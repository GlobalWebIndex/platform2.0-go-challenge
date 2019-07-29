package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// logWebServer a wrapper to log to server
func logWebServer(msg string) {
	log.New(os.Stdout, "http: ", log.LstdFlags).Println(msg)
}

func auth(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, _ := r.BasicAuth()
		if !check(user, pass) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		f(w, r)
	}
}

// timedWrapperHandler a function that act as a wrapper for handlers so that we log each request execution time
func timedWrapperHandleFunc(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f(w, r)
		end := time.Now()
		logWebServer(fmt.Sprintf("Request: %q took: %v", html.EscapeString(r.URL.Path), end.Sub(start)))
	}
}

// ping request a simple ping for health check
func ping(w http.ResponseWriter, r *http.Request) {
	// nothing really here
	w.WriteHeader(http.StatusOK)
}

// addRandomAssets request add random assets, can be helpful
func addRandomAssets(w http.ResponseWriter, r *http.Request) {
	// Get params
	params := mux.Vars(r)

	// check param format
	num, err := strconv.Atoi(params["num"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fillAssets(num)
	w.WriteHeader(http.StatusOK)
}

// getAsset request asset with provided {assetid}
func getAsset(w http.ResponseWriter, r *http.Request) {
	// Get params
	params := mux.Vars(r)

	// check param format
	assetID, err := strconv.Atoi(params["assetid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check asset existence
	asset, found := DBgetAssetByID(assetID)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// set response to json
	w.Header().Set("Content-Type", "application/json")

	// success
	w.WriteHeader(http.StatusOK)

	// convert asset to json format
	json.NewEncoder(w).Encode(asset)

}

// getAssets request retrieves all assets
func getAssets(w http.ResponseWriter, r *http.Request) {
	// retrieve all assets
	allAssets := DBgetAllAssets()

	// set response to json
	w.Header().Set("Content-Type", "application/json")

	// success
	w.WriteHeader(http.StatusOK)

	// assets grouped by based on their type
	aGrp := getAssetsGrouped(allAssets)

	// convert assets json format
	json.NewEncoder(w).Encode(aGrp)
}

// getUserAssets request all assets of a userid
func getUserAssets(w http.ResponseWriter, r *http.Request) {
	// Get params
	params := mux.Vars(r)

	// check param format
	userID, err := strconv.Atoi(params["userid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check user existence
	userAssets, found := DBgetUserAssets(userID)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// set response to json
	w.Header().Set("Content-Type", "application/json")

	// success
	w.WriteHeader(http.StatusOK)

	// assets grouped by based on their type
	aGrp := getAssetsGrouped(userAssets)

	// convert assets json format
	json.NewEncoder(w).Encode(aGrp)
}

// validationChecks auxiliary helper function for similar requests
// returns an asset that is requested to perform actions on it
// in case sth is wrong the asset is nil with the respective http status
func validationChecks(r *http.Request) (int, *Asset) {
	// Get params
	params := mux.Vars(r)

	// check param format
	userID, err := strconv.Atoi(params["userid"])
	if err != nil {
		return http.StatusBadRequest, nil
	}
	// check param format
	assetID, err := strconv.Atoi(params["assetid"])
	if err != nil {
		return http.StatusBadRequest, nil
	}

	// check asset existence
	asset, found := DBgetAssetByID(assetID)
	if !found {
		return http.StatusNotFound, nil
	}

	// check asset ownwership
	if getUserId(asset) != userID {
		return http.StatusNotFound, nil
	}

	return http.StatusOK, asset
}

// removeAsset request remove asset
func removeAsset(w http.ResponseWriter, r *http.Request) {
	// perform all prerequisite checks
	statusCode, asset := validationChecks(r)

	// if asset not returned
	if asset == nil {
		w.WriteHeader(statusCode)
		return
	}

	// success
	w.WriteHeader(http.StatusOK)

	// removal
	DBremoveAsset(asset)
}

// favorAsset request to mark asset as favorite
func favorAsset(w http.ResponseWriter, r *http.Request) {
	// perform all prerequisite checks
	statusCode, asset := validationChecks(r)

	// if asset not returned
	if asset == nil {
		w.WriteHeader(statusCode)
		return
	}
	// set response to json
	w.Header().Set("Content-Type", "application/json")

	// success
	w.WriteHeader(http.StatusOK)

	// mark it as favorite
	FavorAsset(asset)

	// respond it to uset
	json.NewEncoder(w).Encode(asset)
}

// unfavorAsset request unmark asset from favorite
func unfavorAsset(w http.ResponseWriter, r *http.Request) {

	// perform all prerequisite checks
	statusCode, asset := validationChecks(r)

	// if asset not returned
	if asset == nil {
		w.WriteHeader(statusCode)
		return
	}

	// set response to json
	w.Header().Set("Content-Type", "application/json")

	// success
	w.WriteHeader(http.StatusOK)

	// mark it as favorite
	UnFavorAsset(asset)

	// respond it to uset
	json.NewEncoder(w).Encode(asset)
}

func editDescAsset(w http.ResponseWriter, r *http.Request) {

	// perform all prerequisite checks
	statusCode, assetToEdit := validationChecks(r)

	// if asset not returned
	if assetToEdit == nil {
		w.WriteHeader(statusCode)
		return
	}

	// the json expected is just a json formatted as {"newdesc":"newvalueOfDesc"}
	type jsonExpected struct {
		NewDesc string `json:"newdesc"`
	}
	// get the value from the post request with the necessary checks
	jsonReceived := jsonExpected{}
	err := json.NewDecoder(r.Body).Decode(&jsonReceived)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// set response to json
	w.Header().Set("Content-Type", "application/json")

	// success
	w.WriteHeader(http.StatusOK)

	// do the actual edit
	EditDesc(assetToEdit, jsonReceived.NewDesc)

	// respond it to uset
	json.NewEncoder(w).Encode(assetToEdit)
}

// endPointHandler having handler in seperate function helps especially in testing
func endPointHandler() http.Handler {
	r := mux.NewRouter()

	// ping
	r.HandleFunc("/ping", timedWrapperHandleFunc(ping)).Methods("GET")

	// add random assets
	r.HandleFunc("/assets/add/{num}", timedWrapperHandleFunc(auth(addRandomAssets))).Methods("POST")

	// specific id
	r.HandleFunc("/asset/{assetid}", timedWrapperHandleFunc(auth(getAsset))).Methods("GET")

	// all assets
	r.HandleFunc("/assets", timedWrapperHandleFunc(auth(getAssets))).Methods("GET")

	// all assets of user
	r.HandleFunc("/user/{userid}", timedWrapperHandleFunc(auth(getUserAssets))).Methods("GET")

	// specific asset of user
	r.HandleFunc("/user/{userid}/asset/{assetid}", timedWrapperHandleFunc(auth(removeAsset))).Methods("DELETE")

	// mark asset as favorite
	r.HandleFunc("/user/{userid}/asset/{assetid}/favor", timedWrapperHandleFunc(auth(favorAsset))).Methods("PUT")

	// unmark asset as favorite
	r.HandleFunc("/user/{userid}/asset/{assetid}/unfavor", timedWrapperHandleFunc(auth(unfavorAsset))).Methods("PUT")

	// edit asset's description
	r.HandleFunc("/user/{userid}/asset/{assetid}/editdesc", timedWrapperHandleFunc(auth(editDescAsset))).Methods("PUT")

	return r
}
