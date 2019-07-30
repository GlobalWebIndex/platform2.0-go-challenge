package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

	// convert assets to json format
	json.NewEncoder(w).Encode(aGrp)
}

// getUsers request retrieves all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	// retrieve all users
	allUsers := DBgetAllUsers()

	// set response to json
	w.Header().Set("Content-Type", "application/json")

	// success
	w.WriteHeader(http.StatusOK)

	// convert users to json format
	json.NewEncoder(w).Encode(allUsers)
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

	// check if the request is for the same user
	un, _, _ := r.BasicAuth()

	ugt, found := DBgetUserNameByID(userID)
	if !found {
		// theres no such user
		w.WriteHeader(http.StatusNotFound)
	}
	if ugt != un {
		w.WriteHeader(http.StatusUnauthorized)
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
	r.HandleFunc("/ping", ping).Methods("GET")

	// add random assets
	r.HandleFunc("/assets/add/{num}", adminauth(addRandomAssets)).Methods("POST")

	// specific id
	r.HandleFunc("/asset/{assetid}", adminauth(getAsset)).Methods("GET")

	// all assets
	r.HandleFunc("/assets", adminauth(getAssets)).Methods("GET")

	// all users
	r.HandleFunc("/users", adminauth(getUsers)).Methods("GET")

	// all assets of user
	r.HandleFunc("/user/{userid}", userauth(getUserAssets)).Methods("GET")

	// removes specific asset of user
	r.HandleFunc("/user/{userid}/asset/{assetid}", userauth(removeAsset)).Methods("DELETE")

	// mark asset as favorite
	r.HandleFunc("/user/{userid}/asset/{assetid}/favor", userauth(favorAsset)).Methods("PUT")

	// unmark asset as favorite
	r.HandleFunc("/user/{userid}/asset/{assetid}/unfavor", userauth(unfavorAsset)).Methods("PUT")

	// edit asset's description
	r.HandleFunc("/user/{userid}/asset/{assetid}/editdesc", userauth(editDescAsset)).Methods("PUT")

	return r
}
