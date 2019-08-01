package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

const port = 8081

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func getAllAssets(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(assets)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func addToFavorites(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for i := 0; i < len(assets); i++ {
		if assets[i].ID == id {
			assets[i].Favorite = true
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

func deleteFromFavorites(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for i := 0; i < len(assets); i++ {
		if assets[i].ID == id {
			assets[i].Favorite = false
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

func editAsset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader((http.StatusBadRequest))
		return
	}

	var request editAssetRequest
	err = json.Unmarshal(reqBody, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i := 0; i < len(assets); i++ {
		if assets[i].ID == id {
			assets[i].Description = request.Description
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

func handleRequests() error {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/assets", getAllAssets).Methods("GET")
	myRouter.HandleFunc("/asset/favorite/{id}", addToFavorites).Methods("POST")
	myRouter.HandleFunc("/asset/favorite/{id}", deleteFromFavorites).Methods("DELETE")
	myRouter.HandleFunc("/asset/{id}", editAsset).Methods("POST")
	return http.ListenAndServe(fmt.Sprintf(":%d", port), myRouter)
}
