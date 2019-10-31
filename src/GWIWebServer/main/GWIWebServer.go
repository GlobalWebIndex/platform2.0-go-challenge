package main

import (
	"GWIWebServer/eventHandler"
	"fmt"
	"net/http"
)

func getParamValue(param string, w http.ResponseWriter, r *http.Request) (string, bool) {
	params, ok := r.URL.Query()[param]
	if !ok || len(params[0]) < 1 {
		w.Write([]byte("URL Param '" + param + "' is missing\n"))
		return "", ok
	}
	value := params[0]

	if param != "userID" && param != "assetID" && param != "description" {
		w.Write([]byte("Invalid Parameter '" + param + "'\n"))
		ok = false
	}
	return value, ok
}

func showUsersRequest(w http.ResponseWriter, r *http.Request) {
	out := eventHandler.ShowUsers()
	w.Write([]byte(out))
}

func showAssetsRequest(w http.ResponseWriter, r *http.Request) {
	out := eventHandler.ShowAssets()
	w.Write([]byte(out))
}

func showFavoritesRequest(w http.ResponseWriter, r *http.Request) {
	userID, okUserID := getParamValue("userID", w, r)

	if okUserID {
		out := eventHandler.ShowFavorites(userID)
		w.Write([]byte(out))
	}
}

func addFavoriteRequest(w http.ResponseWriter, r *http.Request) {
	userID, okUserID := getParamValue("userID", w, r)
	assetID, okAssetID := getParamValue("assetID", w, r)

	if okUserID && okAssetID {
		out := eventHandler.AddFavorite(userID, assetID)
		w.Write([]byte(out))
	}
}

func deleteFavoriteRequest(w http.ResponseWriter, r *http.Request) {
	userID, okUserID := getParamValue("userID", w, r)
	assetID, okAssetID := getParamValue("assetID", w, r)

	if okUserID && okAssetID {
		out := eventHandler.DeleteFavorite(userID, assetID)
		w.Write([]byte(out))
	}
}

func editAssetDescriptionRequest(w http.ResponseWriter, r *http.Request) {
	assetID, okAssetID := getParamValue("assetID", w, r)
	description, okDescription := getParamValue("description", w, r)

	if okAssetID && okDescription {
		out := eventHandler.EditAssetDescription(assetID, description)
		w.Write([]byte(out))
	}
}

func runSim() {
	usersNumber := 25000
	assetsNumber := 50000

	for i := 0; i < usersNumber; i++ {
		username := "User_" + fmt.Sprint(i)
		eventHandler.AddNewUser(username)
	}

	for i := 0; usersNumber > 0 && i < assetsNumber; i++ {
		var assetID string
		if i%3 == 0 {
			assetID = "Chart_" + fmt.Sprint(i)
			eventHandler.AddNewChart(assetID, "", "", "", "")
		} else if i%3 == 1 {
			assetID = "Insight_" + fmt.Sprint(i)
			eventHandler.AddNewInsight(assetID, "", "")
		} else {
			assetID = "Audience_" + fmt.Sprint(i)
			eventHandler.AddNewAudience(assetID, "", 0, "", 0, 0, 0)
		}

		username := "User_" + fmt.Sprint(i%usersNumber)
		eventHandler.AddFavorite(username, assetID)
	}
}

func main() {
	runSim()

	http.HandleFunc("/GWI-platform/show-users/", showUsersRequest)
	http.HandleFunc("/GWI-platform/show-assets/", showAssetsRequest)
	http.HandleFunc("/GWI-platform/show-favorites/", showFavoritesRequest)
	http.HandleFunc("/GWI-platform/add-favorite/", addFavoriteRequest)
	http.HandleFunc("/GWI-platform/delete-favorite/", deleteFavoriteRequest)
	http.HandleFunc("/GWI-platform/edit-asset-description/", editAssetDescriptionRequest)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
