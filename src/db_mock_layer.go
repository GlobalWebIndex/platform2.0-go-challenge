package main

// this file plays the role DB
// if there was a DB, only this file has to change keeping concerns seperate

// map of assets based on the userId, so that the assets of a given user is retrieved reasonably quick
var assetsByUser = make(map[int]map[int]*Asset)

// map of assets based on the assetId, so
var assetByID = make(map[int]*Asset)

// a dummy var that stores all the credentials
var allcreds = make(map[string]string)

// DBaddCreds just create some admin users here
func DBaddCreds() {
	allcreds["vasilis"] = "12345"
}

// DBaddAsset a function that assigns the reference of asset to the internal maps
// DB equivalent : `insert into tasset(...) values (...)`
func DBaddAsset(a *Asset) {

	// get user index
	userID := getUserId(a)
	// get asset index
	assetID := getAssetId(a)

	// init map in case first asset for that user
	if assetsByUser[userID] == nil {
		assetsByUser[userID] = make(map[int]*Asset)
	}

	// assign the pointer
	assetsByUser[userID][assetID] = a
	assetByID[assetID] = a
}

// DBremoveAsset a function that removes the asset from the internal maps
// DB equivalent : `delete from tasset where assetID = ?`
func DBremoveAsset(a *Asset) {
	// get the userId, before removing anything
	// get user index
	userID := getUserId(a)
	// get asset index
	assetID := getAssetId(a)

	// remove from map [assetId]
	delete(assetByID, assetID)

	// remove from map [userId][assetId]

	// delete the pointer from the user
	delete(assetsByUser[userID], assetID)

	//in case there are no assets left for the user also delete the map
	if len(assetsByUser[userID]) == 0 {
		delete(assetsByUser, userID)
	}
}

// DBgetUserAssets a function that returns all the assets of a user
// DB equivalent: `select * from tassets where user_id = ?`
func DBgetUserAssets(userID int) (map[int]*Asset, bool) {
	mapValue, found := assetsByUser[userID]
	return mapValue, found
}

// DBgetAllAssets a function that returns all assets, mainly for testing
// DB equivalent: `select * from tassets`
func DBgetAllAssets() map[int]*Asset {
	// just return all the map
	return assetByID
}

// DBgetAssetByID a fuction that returns an asset given its id, and whether it was found
// DB equivalent: `select * from tassets where assetId = ?`
func DBgetAssetByID(assetID int) (*Asset, bool) {
	value, found := assetByID[assetID]
	return value, found
}

// DBupdateAssetPersist a dummy function, since everything is on memory, changes are persisted on the fly
// DB equivalent: `update tasset ....`
func DBupdateAssetPersist(a *Asset) {
	// nothing atm
}

// DBgetUserSL a function that retrieves the user ids into a slice of int
// DB equivalent: `select userid from ...`
func DBgetUserSL() []int {
	uids := []int{}
	for k := range assetsByUser {
		uids = append(uids, k)
	}
	return uids
}

// DBgetAsetIdsSL a funtion that retrieves the asset ids into a slice of int
// DB equivalent: `select assetid from ...`
func DBgetAsetIdsSL() []int {
	aids := []int{}
	for k := range assetByID {
		aids = append(aids, k)
	}
	return aids
}

// DBgetAssetsCount a funtion that count the number of assets
// DB equivalent `select count(*) from ...`
func DBgetAssetsCount() int {
	return len(assetByID)
}

// removeAllAssets a funtion that removes every asset
// DB equivalent `delete tasset ...`
func removeAllAssets() {

	assetIds := DBgetAsetIdsSL()
	for _, aID := range assetIds {
		a, found := assetByID[aID]
		if found {
			DBremoveAsset(a)
		}
	}
}
