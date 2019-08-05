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
func (db *dbMock) addRandomAssets(w http.ResponseWriter, r *http.Request) {
	// Get params
	params := mux.Vars(r)

	// check param format
	num, err := strconv.Atoi(params["num"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	db.fillAssets(num)
	w.WriteHeader(http.StatusOK)
}

// getAsset request asset with provided {assetid}
func (db *dbMock) getAsset(w http.ResponseWriter, r *http.Request) {
	// Get params
	params := mux.Vars(r)

	// check param format
	assetID, err := strconv.Atoi(params["assetid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check asset existence
	asset, found := db.DBgetAssetByID(assetID)
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
func (db *dbMock) getAssets(w http.ResponseWriter, r *http.Request) {
	// retrieve all assets
	allAssets := db.DBgetAllAssets()

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
func (db *dbMock) getUsers(w http.ResponseWriter, r *http.Request) {
	// retrieve all users
	allUsers := db.DBgetAllUsers()

	// set response to json
	w.Header().Set("Content-Type", "application/json")

	// success
	w.WriteHeader(http.StatusOK)

	// convert users to json format
	json.NewEncoder(w).Encode(allUsers)
}

// getUserAssets request all assets of a userid
func (db *dbMock) getUserAssets(w http.ResponseWriter, r *http.Request) {
	// Get params
	params := mux.Vars(r)

	// check param format
	userID, err := strconv.Atoi(params["userid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check user existence
	userAssets, found := db.DBgetUserAssets(userID)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// check if the request is for the same user
	un, _, _ := r.BasicAuth()

	ugt, found := db.DBgetUserNameByID(userID)
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
func (db *dbMock) removeAsset(w http.ResponseWriter, r *http.Request) {
	// perform all prerequisite checks

	statusCode, asset := db.validationChecks(r)

	// if asset not returned
	if asset == nil {
		w.WriteHeader(statusCode)
		return
	}

	// success
	w.WriteHeader(http.StatusOK)

	// removal
	db.DBremoveAsset(asset)
}

// favorAsset request to mark asset as favorite
func (db *dbMock) favorAsset(w http.ResponseWriter, r *http.Request) {
	// perform all prerequisite checks
	statusCode, asset := db.validationChecks(r)

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
	db.FavorAsset(asset)

	// respond it to uset
	json.NewEncoder(w).Encode(asset)
}

// unfavorAsset request unmark asset from favorite
func (db *dbMock) unfavorAsset(w http.ResponseWriter, r *http.Request) {

	// perform all prerequisite checks
	statusCode, asset := db.validationChecks(r)

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
	db.UnFavorAsset(asset)

	// respond it to uset
	json.NewEncoder(w).Encode(asset)
}

func (db *dbMock) editDescAsset(w http.ResponseWriter, r *http.Request) {

	// perform all prerequisite checks
	statusCode, assetToEdit := db.validationChecks(r)

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
	db.EditDesc(assetToEdit, jsonReceived.NewDesc)

	// respond it to uset
	json.NewEncoder(w).Encode(assetToEdit)
}

// endPointHandler having handler in seperate function helps especially in testing
func endPointHandler(db *dbMock) http.Handler {
	r := mux.NewRouter()

	// ping
	r.HandleFunc("/ping", ping).Methods("GET")

	// add random assets
	r.HandleFunc("/assets/add/{num}", db.adminauth(db.addRandomAssets)).Methods("POST")

	// specific id
	r.HandleFunc("/asset/{assetid}", db.adminauth(db.getAsset)).Methods("GET")

	// all assets
	r.HandleFunc("/assets", db.adminauth(db.getAssets)).Methods("GET")

	// all users
	r.HandleFunc("/users", db.adminauth(db.getUsers)).Methods("GET")

	// all assets of user
	r.HandleFunc("/user/{userid}", db.userauth(db.getUserAssets)).Methods("GET")

	// removes specific asset of user
	r.HandleFunc("/user/{userid}/asset/{assetid}", db.userauth(db.removeAsset)).Methods("DELETE")

	// mark asset as favorite
	r.HandleFunc("/user/{userid}/asset/{assetid}/favor", db.userauth(db.favorAsset)).Methods("PUT")

	// unmark asset as favorite
	r.HandleFunc("/user/{userid}/asset/{assetid}/unfavor", db.userauth(db.unfavorAsset)).Methods("PUT")

	// edit asset's description
	r.HandleFunc("/user/{userid}/asset/{assetid}/editdesc", db.userauth(db.editDescAsset)).Methods("PUT")

	return r
}
