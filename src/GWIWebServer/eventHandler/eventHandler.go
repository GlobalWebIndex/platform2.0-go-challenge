package eventHandler

import (
	"GWIWebServer/assets"
	"GWIWebServer/users"
	"bytes"
	"fmt"
)

var userIDMap = make(map[string]users.User)
var assetIDMap = make(map[string]assets.AssetInterface)

func checkUserExists(userID string) bool {
	_, ok := userIDMap[userID]
	return ok
}

func checkAssetExists(assetID string) bool {
	_, ok := assetIDMap[assetID]
	return ok
}

// AddNewUser adds an new user to database
func AddNewUser(userID string) {
	if _, ok := userIDMap[userID]; !ok {
		userIDMap[userID] = users.CreateNewUser(userID)
	} else {
		fmt.Println("UserID '" + userID + "' already exists")
	}
}

// AddNewChart adds an new chart to database
func AddNewChart(assetID string, description string, title string, xAxisTitle string, yAxisTitle string) {
	if _, ok := assetIDMap[assetID]; !ok {
		assetIDMap[assetID] = assets.CreateNewChart(assetID, description, title, xAxisTitle, yAxisTitle)
	} else {
		fmt.Println("AssetID '" + assetID + "' already exists")
	}
}

// AddNewInsight adds an new insight to database
func AddNewInsight(assetID string, description string, text string) {
	if _, ok := assetIDMap[assetID]; !ok {
		assetIDMap[assetID] = assets.CreateNewInsight(assetID, description, text)
	} else {
		fmt.Println("AssetID '" + assetID + "' already exists")
	}
}

// AddNewAudience adds an new audience to database
func AddNewAudience(assetID string, description string, gender assets.Gender, birthCountry string, ageGroup assets.AgeGroup, hoursOnSocialMedia uint8, lastMonthPurchases uint32) {
	if _, ok := assetIDMap[assetID]; !ok {
		assetIDMap[assetID] = assets.CreateNewAudience(assetID, description, gender, birthCountry, ageGroup, hoursOnSocialMedia, lastMonthPurchases)
	} else {
		fmt.Println("AssetID '" + assetID + "' already exists")
	}
}

// ShowUsers returns the list of all users
func ShowUsers() string {
	var out bytes.Buffer
	out.WriteString("Users List: (" + fmt.Sprint(len(userIDMap)) + ")\n\n")
	for k := range userIDMap {
		out.WriteString(k + "\n")
	}
	return out.String()
}

// ShowAssets returns the list of all assets
func ShowAssets() string {
	var out bytes.Buffer
	out.WriteString("Assets List: (" + fmt.Sprint(len(assetIDMap)) + ")\n\n")
	for _, v := range assetIDMap {
		out.WriteString(v.GetDetails())
	}
	return out.String()
}

// ShowFavorites returns the favorite list of a specific user
func ShowFavorites(userID string) string {
	var out bytes.Buffer
	userExists := checkUserExists(userID)

	if userExists {
		user := userIDMap[userID]
		userFavorites := user.GetUserFavorites()

		out.WriteString("Favorites for UserID '" + userID + "': (" + fmt.Sprint(len(userFavorites)) + ")\n\n")
		for k := range userFavorites {
			out.WriteString(assetIDMap[k].GetDetails())
		}
	} else {
		out.WriteString("UserID '" + userID + "' not found\n")
	}

	return out.String()
}

// AddFavorite adds an asset in favorite list of a specific user
func AddFavorite(userID string, assetID string) string {
	var out bytes.Buffer
	userExists := checkUserExists(userID)
	assetExists := checkAssetExists(assetID)

	if userExists && assetExists {
		user := userIDMap[userID]
		userFavorites := user.GetUserFavorites()

		if _, ok := userFavorites[assetID]; !ok {
			userFavorites[assetID] = struct{}{}
			out.WriteString("Added AssetID '" + assetID + "' to favorites for UserID '" + userID + "'\n")
		} else {
			out.WriteString("AssetID '" + assetID + "' already in favorites for UserID '" + userID + "'\n")
		}
	} else {
		if !userExists {
			out.WriteString("UserID '" + userID + "' not found\n")
		}
		if !assetExists {
			out.WriteString("AssetID '" + assetID + "' not found\n")
		}
	}

	return out.String()
}

// DeleteFavorite removes an asset from the favorite list of a specific user
func DeleteFavorite(userID string, assetID string) string {
	var out bytes.Buffer
	userExists := checkUserExists(userID)
	assetExists := checkAssetExists(assetID)

	if userExists && assetExists {
		user := userIDMap[userID]
		userFavorites := user.GetUserFavorites()

		if _, ok := userFavorites[assetID]; ok {
			delete(userFavorites, assetID)
			out.WriteString("Deleted AssetID '" + assetID + "' from favorites of UserID '" + userID + "'\n")
		} else {
			out.WriteString("AssetID '" + assetID + "' not in favorites for UserID '" + userID + "'\n")
		}
	} else {
		if !userExists {
			out.WriteString("UserID '" + userID + "' not found\n")
		}
		if !assetExists {
			out.WriteString("AssetID '" + assetID + "' not found\n")
		}
	}

	return out.String()
}

// EditAssetDescription edits the description of an asset
func EditAssetDescription(assetID string, description string) string {
	var out bytes.Buffer
	assetExists := checkAssetExists(assetID)

	if assetExists {
		switch assetIDMap[assetID].GetAsset().AssetType {
		case assets.ChartType:
			chart := assetIDMap[assetID].(assets.Chart)
			chart.Asset.SetDescription(description)
			assetIDMap[assetID] = chart
		case assets.InsightType:
			insight := assetIDMap[assetID].(assets.Insight)
			insight.Asset.SetDescription(description)
			assetIDMap[assetID] = insight
		case assets.AudienceType:
			audience := assetIDMap[assetID].(assets.Audience)
			audience.Asset.SetDescription(description)
			assetIDMap[assetID] = audience
		}

		out.WriteString("New Description for AssetID '" + assetID + "': '" + assetIDMap[assetID].GetAsset().GetDescription() + "'\n")
	} else {
		out.WriteString("AssetID '" + assetID + "' not found\n")
	}

	return out.String()
}
