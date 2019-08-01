package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetAllAssetsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/assets", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAllAssets)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestAddToFavoritesHandlerCorrect(t *testing.T) {
	req, err := http.NewRequest("POST", "/asset/favorite/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/asset/favorite/{id}", addToFavorites).Methods("POST")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	for i := 0; i < len(assets); i++ {
		if assets[i].ID == "2" {
			if !assets[i].Favorite {
				t.Fatal("Asset with ID 2 should have favorite set to true")
			}
			return
		}
	}
	t.Fatalf("Asset with ID 2 not found")
}

func TestAddToFavoritesHandlerWrong(t *testing.T) {
	req, err := http.NewRequest("POST", "/asset/favorite/sasasa", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/asset/favorite/{id}", addToFavorites).Methods("POST")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func TestDeleteToFavoritesHandler(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/asset/favorite/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/asset/favorite/{id}", deleteFromFavorites).Methods("DELETE")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	for i := 0; i < len(assets); i++ {
		if assets[i].ID == "1" {
			if assets[i].Favorite {
				t.Fatal("Asset with ID 1 should have favorite set to false")
			}
			return
		}
	}
	t.Fatalf("Asset with ID 1 not found")
}

func TestEditAssetHandler(t *testing.T) {

	newAsset := asset{
		ID:          "1",
		Description: "newdescription",
	}

	b, err := json.Marshal(newAsset)

	if err != nil {
		t.Fatal("Cannot convert to JSON")
	}

	req, err := http.NewRequest("POST", "/asset/1", bytes.NewBuffer(b))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/asset/{id}", editAsset).Methods("POST")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	for i := 0; i < len(assets); i++ {
		if assets[i].ID == "1" {
			if assets[i].Description != newAsset.Description {
				t.Fatalf("Asset with ID 1 should have description set to %s", newAsset.Description)
			}
			return
		}
	}
	t.Fatalf("Asset with ID 1 not found")
}
