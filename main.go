package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

//User struct
type User struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Asset []Asset `json:"asset"`
}

//Asset struct
type Asset struct {
	ID          string            `json:"id"`
	Description string            `json:"description"`
	AssetType   string            `json:"assettype"`
	Data        map[string]string `json:"data"`
	IsFav       bool              `json:"isfav"`
}

//Init structs
var users []User
var assets []Asset

func main() {
	r := mux.NewRouter()
	r.Use(simpleMw)
	//mock data
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}
	var data1 = map[string]string{
		"xtitle": "x",
		"ytitle": "y",
		"xdata":  "3",
		"ydata":  "4",
	}

	var data2 = map[string]string{
		"gender":             "male",
		"country":            "greece",
		"ageGroups":          "25-34",
		"hoursOnSocial":      "40",
		"purchasesLastMonth": "33",
	}

	initassets := []Asset{
		Asset{
			ID:          "1",
			Description: "chart1",
			AssetType:   "chart",
			IsFav:       true,
			Data:        data1,
		},
		Asset{
			ID:          "2",
			Description: "audience1",
			AssetType:   "audience",
			IsFav:       true,
			Data:        data2,
		},
	}

	users = append(users, User{ID: "1", Name: "Tsolakidis",
		Asset: initassets,
	})
	assets = initassets

	r.HandleFunc("/api/health", health).Methods("GET")
	r.HandleFunc("/api/user/", createUser).Methods("POST")
	r.HandleFunc("/api/user/{id}", getUser).Methods("GET")
	r.HandleFunc("/api/asset/{userid}", createAsset).Methods("PUT")
	r.HandleFunc("/api/assets", getAssets).Methods("GET")
	r.HandleFunc("/api/user/{userid}/asset/{id}", updateAsset).Methods("PUT")
	r.HandleFunc("/api/user/{userid}/asset/{id}/fav", favAsset).Methods("PUT")

	var port string

	port = os.Getenv("PORT")

	fmt.Printf("Web server started on port " + port + "\n")

	log.Fatal(http.ListenAndServe(":"+port, r))

}
func simpleMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "App is running \n")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application-json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(1000000))
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application-json")
	params := mux.Vars(r)
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func createAsset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application-json")
	var asset Asset
	_ = json.NewDecoder(r.Body).Decode(&asset)
	asset.ID = strconv.Itoa(rand.Intn(1000000))
	params := mux.Vars(r)
	for i, item := range users {
		if item.ID == params["userid"] {

			assets = append(item.Asset, asset)

			users[i].Asset = assets
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}
	json.NewEncoder(w).Encode(assets)
}
func getAssets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application-json")
	json.NewEncoder(w).Encode(assets)
}

func updateAsset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application-json")
	var asset Asset
	var newAsset Asset
	_ = json.NewDecoder(r.Body).Decode(&newAsset)
	params := mux.Vars(r)

	for j, item := range users {
		if item.ID == params["userid"] {
			for i, sitem := range item.Asset {
				if sitem.ID == params["id"] {
					asset = sitem
					asset.Description = newAsset.Description

					assets = append(assets[:i], assets[i+1:]...)
					assets = append(assets, asset)
					users[j].Asset = assets
					json.NewEncoder(w).Encode(users[j])
					return
				}
			}
		}
	}

	json.NewEncoder(w).Encode(assets)
}

func favAsset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application-json")
	var asset Asset
	var newAsset Asset
	_ = json.NewDecoder(r.Body).Decode(&newAsset)
	params := mux.Vars(r)

	for j, item := range users {
		if item.ID == params["userid"] {
			for i, sitem := range item.Asset {
				if sitem.ID == params["id"] {
					asset = sitem
					asset.IsFav = newAsset.IsFav

					assets = append(assets[:i], assets[i+1:]...)
					assets = append(assets, asset)
					users[j].Asset = assets
					json.NewEncoder(w).Encode(users[j])
					return
				}
			}
		}
	}

	json.NewEncoder(w).Encode(assets)
}
