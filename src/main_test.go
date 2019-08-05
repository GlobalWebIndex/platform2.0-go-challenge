package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

// a simple ping test to the webserver just to verify its listening
func TestEndPointPing(t *testing.T) {
	// create a new DB
	db := CreateNewDB()

	srv := httptest.NewServer(endPointHandler(db))
	defer srv.Close()
	url := fmt.Sprintf("%s/ping", srv.URL)
	res, err := http.Get(url)

	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", res.StatusCode)
	}
}

func TestEndPointasset(t *testing.T) {
	// create a new DB
	db := CreateNewDB()

	// add a single asset
	db.fillAssets(2)
	db.DBaddCreds()

	srv := httptest.NewServer(endPointHandler(db))
	defer srv.Close()

	// get some existing asset ids
	realAssetIds := db.DBgetAsetIdsSL()

	tt := []struct {
		testName       string
		endPoint       string
		sp1            string
		userid         string
		username       string
		password       string
		statusExpected int
	}{
		{"Wrong EndPoint:", "/asset", "/", "", "admin", "12345", http.StatusNotFound},
		{"Wrong EndPoint without var:", "/asset/", "", "", "admin", "12345", http.StatusNotFound},
		{"Wrong EndPoint with invalid assetId:", "/asset/asd", "", "", "admin", "12345", http.StatusBadRequest},
		{"Non-Existing asset0:", "/asset", "", "10000", "admin", "12345", http.StatusNotFound},
		{"This should not be allowed :", "/asset/1/asd", "", strconv.Itoa(0), "admin", "12345", http.StatusNotFound},
		{"Existing asset1:", "/asset", "/", strconv.Itoa(realAssetIds[0]), "admin", "12345", http.StatusOK},
		{"Existing asset2:", "/asset", "/", strconv.Itoa(realAssetIds[1]), "admin", "12345", http.StatusOK},
		{"Existing asset3 but unauthorized:", "/asset", "/", strconv.Itoa(realAssetIds[1]), "admin", "123", http.StatusUnauthorized},
		{"Existing asset3 but unauthorized 2:", "/asset", "/", strconv.Itoa(realAssetIds[1]), "vas", "12345", http.StatusUnauthorized},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {

			wholeURL := srv.URL + tc.endPoint + tc.sp1 + tc.userid

			client := &http.Client{}
			req, err := http.NewRequest("GET", wholeURL, nil)
			if err != nil {
				t.Fatalf("could not send GET request: %v", err)
			}
			req.SetBasicAuth(tc.username, tc.password)
			res, err := client.Do(req)
			if err != nil {
				t.Fatalf("Error: %v", err)
			}

			// final check to verify correctness
			if res.StatusCode != tc.statusExpected {
				t.Errorf("%v: Expected status %v got %v", wholeURL, tc.statusExpected, res.StatusCode)
			}

		})
	}
	db.removeAllAssets()
}

func TestEndgetUserAssets(t *testing.T) {
	// create a new DB
	db := CreateNewDB()

	// add a single asset
	db.fillAssets(3)
	db.DBaddCreds()

	srv := httptest.NewServer(endPointHandler(db))
	defer srv.Close()

	// get some existing user_ids
	realUserIds := db.DBgetUserSL()

	gtu, _ := db.DBgetUserNameByID(realUserIds[0])
	gtp, _ := db.DBgetUserPassByID(realUserIds[0])

	tt := []struct {
		testName       string
		endPoint       string
		sp1            string
		userid         string
		username       string
		password       string
		statusExpected int
	}{
		{"Wrong EndPoint without real user", "/user/ss", "", strconv.Itoa(0), gtu, gtp, http.StatusBadRequest},
		{"Non Existing user ", "/user", "/", strconv.Itoa(900000), gtu, gtp, http.StatusNotFound},
		{"Existing user 1", "/user", "/", strconv.Itoa(realUserIds[0]), gtu, gtp, http.StatusOK},
		{"Existing user 2", "/user", "/", strconv.Itoa(realUserIds[0]), gtu, gtp, http.StatusOK},
		{"Existing user 3", "/user", "/", strconv.Itoa(realUserIds[0]), gtu, gtp, http.StatusOK},
		{"User not same as auth", "/user", "/", strconv.Itoa(realUserIds[1]), gtu, gtp, http.StatusUnauthorized},
		{"Existing user 3 but unauthorized", "/user", "/", strconv.Itoa(realUserIds[0]), "admin", "123", http.StatusUnauthorized},
		{"Existing user 3 but unauthorized 2", "/user", "/", strconv.Itoa(realUserIds[0]), "vasi", "12345", http.StatusUnauthorized},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {

			wholeURL := srv.URL + tc.endPoint + tc.sp1 + tc.userid

			client := &http.Client{}
			req, err := http.NewRequest("GET", wholeURL, nil)
			if err != nil {
				t.Fatalf("could not send GET request: %v", err)
			}
			req.SetBasicAuth(tc.username, tc.password)
			res, err := client.Do(req)
			if err != nil {
				t.Fatalf("Error: %v", err)
			}

			if res.StatusCode != tc.statusExpected {
				t.Errorf(" %v : Expected status %v got %v", wholeURL, tc.statusExpected, res.StatusCode)
			}

		})
	}
	db.removeAllAssets()
}

func TestEndremoveAsset(t *testing.T) {
	// create a new DB
	db := CreateNewDB()

	// add some assets
	db.fillAssets(30)
	db.DBaddCreds()

	srv := httptest.NewServer(endPointHandler(db))
	defer srv.Close()

	// get some existing user_ids
	realUserIds := db.DBgetUserSL()

	// get some existing asset ids
	realAssetIds := db.DBgetAsetIdsSL()

	// get user of asset
	asset, _ := db.DBgetAssetByID(realAssetIds[0])

	userIdOfAsset := getUserId(asset)

	gtu, _ := db.DBgetUserNameByID(realUserIds[1])
	gtp, _ := db.DBgetUserPassByID(realUserIds[1])

	gtuc, _ := db.DBgetUserNameByID(userIdOfAsset)
	gtpc, _ := db.DBgetUserPassByID(userIdOfAsset)

	tt := []struct {
		testName       string
		endPoint       string
		userid         string
		endPoint2      string
		assetid        string
		username       string
		password       string
		statusExpected int
	}{
		{"Non Existent User and non Existent Asset", "/user/", "100000", "/asset/", "100000", gtu, gtp, http.StatusNotFound},
		{"Non Existent User and Existent Asset", "/user/", "10000", "/asset/", strconv.Itoa(realAssetIds[0]), gtu, gtp, http.StatusNotFound},
		{"Existent User and non Existent Asset", "/user/", strconv.Itoa(realUserIds[0]), "/asset/", "10000", gtu, gtp, http.StatusNotFound},
		{"Existent User and Existent Asset but wrong user-to-asset", "/user/", strconv.Itoa(userIdOfAsset + 1), "asset/", strconv.Itoa(realAssetIds[0]), gtu, gtp, http.StatusNotFound},
		{"User not same as auth", "/user/", strconv.Itoa(userIdOfAsset), "/asset/", strconv.Itoa(realAssetIds[0]), gtu, gtp, http.StatusUnauthorized},
		{"Existent User and Existent Asset but unauthorized", "/user/", strconv.Itoa(userIdOfAsset), "/asset/", strconv.Itoa(realAssetIds[0]), "admin", "123", http.StatusUnauthorized},
		{"Existent User and Existent Asset but unauthorized 2", "/user/", strconv.Itoa(userIdOfAsset), "/asset/", strconv.Itoa(realAssetIds[0]), "vas", "12345", http.StatusUnauthorized},
		{"Existent User and Existent Asset", "/user/", strconv.Itoa(userIdOfAsset), "/asset/", strconv.Itoa(realAssetIds[0]), gtuc, gtpc, http.StatusOK},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {

			wholeURL := srv.URL + tc.endPoint + tc.userid + tc.endPoint2 + tc.assetid

			client := &http.Client{}
			req, err := http.NewRequest("DELETE", wholeURL, nil)
			if err != nil {
				t.Fatalf("could not send DELETE request: %v", err)
			}
			req.SetBasicAuth(tc.username, tc.password)
			res, err := client.Do(req)
			if err != nil {
				t.Fatalf("Error: %v", err)
			}

			if res.StatusCode != tc.statusExpected {
				t.Errorf(" %v : Expected status %v got %v", wholeURL, tc.statusExpected, res.StatusCode)
			}

		})
	}
	db.removeAllAssets()
}

func TestEndfavorOrNotAsset(t *testing.T) {
	// create a new DB
	db := CreateNewDB()

	// add some assets
	db.fillAssets(30)
	db.DBaddCreds()

	srv := httptest.NewServer(endPointHandler(db))
	defer srv.Close()

	// get some existing user_ids
	realUserIds := db.DBgetUserSL()

	// get some existing asset ids
	realAssetIds := db.DBgetAsetIdsSL()

	// get user of asset
	asset, _ := db.DBgetAssetByID(realAssetIds[0])

	userIdOfAsset := getUserId(asset)

	gtu, _ := db.DBgetUserNameByID(realUserIds[1])
	gtp, _ := db.DBgetUserPassByID(realUserIds[1])

	gtuc, _ := db.DBgetUserNameByID(userIdOfAsset)
	gtpc, _ := db.DBgetUserPassByID(userIdOfAsset)

	tt := []struct {
		testName       string
		endPoint       string
		userid         string
		endPoint2      string
		assetid        string
		action         string
		username       string
		password       string
		statusExpected int
	}{
		{"favor Non Existent User and non Existent Asset", "/user/", "10000", "/asset/", "10000", "/favor", gtu, gtp, http.StatusNotFound},
		{"favor Non Existent User and Existent Asset", "/user/", "10000", "/asset/", strconv.Itoa(realAssetIds[0]), "/favor", gtu, gtp, http.StatusNotFound},
		{"favor Existent User and non Existent Asset", "/user/", strconv.Itoa(realUserIds[0]), "/asset/", "10000", "/favor", gtu, gtp, http.StatusNotFound},
		{"favor Existent User and Existent Asset but wrong user-to-asset", "/user/", strconv.Itoa(userIdOfAsset + 1), "/asset/", strconv.Itoa(realAssetIds[0]), "/favor", gtu, gtp, http.StatusNotFound},
		{"favor User not same as auth", "/user/", strconv.Itoa(userIdOfAsset), "/asset/", strconv.Itoa(realAssetIds[0]), "/favor", gtu, gtp, http.StatusUnauthorized},
		{"favor Existent User and Existent Asset but unauthorized", "/user/", strconv.Itoa(userIdOfAsset), "/asset/", strconv.Itoa(realAssetIds[0]), "/favor", "admin", "123", http.StatusUnauthorized},
		{"favor Existent User and Existent Asset but unauthorized 2", "/user/", strconv.Itoa(userIdOfAsset), "/asset/", strconv.Itoa(realAssetIds[0]), "/favor", "vas", "12345", http.StatusUnauthorized},
		{"favor Existent User and Existent Asset", "/user/", strconv.Itoa(userIdOfAsset), "/asset/", strconv.Itoa(realAssetIds[0]), "/favor", gtuc, gtpc, http.StatusOK},

		{"unfavor  Non Existent User and non Existent Asset", "/user/", "10000", "/asset/", "10000", "/unfavor", gtu, gtp, http.StatusNotFound},
		{"unfavor Non Existent User and Existent Asset", "/user/", "10000", "/asset/", strconv.Itoa(realAssetIds[0]), "/unfavor", gtu, gtp, http.StatusNotFound},
		{"unfavor Existent User and non Existent Asset", "/user/", strconv.Itoa(realUserIds[0]), "/asset/", "10000", "/unfavor", gtu, gtp, http.StatusNotFound},
		{"unfavor Existent User and Existent Asset but wrong user-to-asset", "/user/", strconv.Itoa(userIdOfAsset + 1), "/asset/", strconv.Itoa(realAssetIds[0]), "/unfavor", gtu, gtp, http.StatusNotFound},
		{"unfavor User not same as auth", "/user/", strconv.Itoa(userIdOfAsset), "/asset/", strconv.Itoa(realAssetIds[0]), "/unfavor", gtu, gtp, http.StatusUnauthorized},
		{"unfavor Existent User and Existent Asset but unauthorized", "/user/", strconv.Itoa(userIdOfAsset), "/asset/", strconv.Itoa(realAssetIds[0]), "/unfavor", "admin", "123", http.StatusUnauthorized},
		{"unfavor Existent User and Existent Asset but unauthorized 2", "/user/", strconv.Itoa(userIdOfAsset), "/asset/", strconv.Itoa(realAssetIds[0]), "/unfavor", "vas", "12345", http.StatusUnauthorized},
		{"unfavor Existent User and Existent Asset", "/user/", strconv.Itoa(userIdOfAsset), "/asset/", strconv.Itoa(realAssetIds[0]), "/unfavor", gtuc, gtpc, http.StatusOK},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {

			wholeURL := srv.URL + tc.endPoint + tc.userid + tc.endPoint2 + tc.assetid + tc.action

			// Create client
			client := &http.Client{}

			// res, err := http.Delete(wholeURL)
			req, err := http.NewRequest("PUT", wholeURL, nil)
			if err != nil {
				t.Fatalf("could not send PUT request: %v", err)
			}
			req.SetBasicAuth(tc.username, tc.password)
			res, err := client.Do(req)
			if err != nil {
				t.Fatalf("Error: %v", err)
			}

			if res.StatusCode != tc.statusExpected {
				t.Errorf(" %v : Expected status %v got %v", wholeURL, tc.statusExpected, res.StatusCode)
			}
		})
	}
	db.removeAllAssets()
}

func TestEndEditDescAsset(t *testing.T) {
	// create a new DB
	db := CreateNewDB()

	// add some assets
	db.fillAssets(30)
	db.DBaddCreds()
	srv := httptest.NewServer(endPointHandler(db))
	defer srv.Close()

	// get some existing user_ids
	realUserIds := db.DBgetUserSL()

	// get some existing asset ids
	realAssetIds := db.DBgetAsetIdsSL()

	// get user of asset
	asset, _ := db.DBgetAssetByID(realAssetIds[0])

	userIdOfAsset := getUserId(asset)

	gtu, _ := db.DBgetUserNameByID(realUserIds[1])
	gtp, _ := db.DBgetUserPassByID(realUserIds[1])

	gtuc, _ := db.DBgetUserNameByID(userIdOfAsset)
	gtpc, _ := db.DBgetUserPassByID(userIdOfAsset)

	tt := []struct {
		testName       string
		endPoint       string
		userid         string
		endPoint2      string
		assetid        string
		username       string
		password       string
		statusExpected int
	}{
		{"Non Existent User and non Existent Asset", "/user/", "10000", "/asset/", "10000", gtu, gtp, http.StatusNotFound},
		{"Non Existent User and Existent Asset", "/user/", "10000", "/asset/", strconv.Itoa(realAssetIds[0]), gtu, gtp, http.StatusNotFound},
		{"Existent User and non Existent Asset", "/user/", strconv.Itoa(realUserIds[0]), "/asset/", "10000", gtu, gtp, http.StatusNotFound},
		{"Existent User and Existent Asset but wrong user-to-asset", "/user/", strconv.Itoa(userIdOfAsset + 1), "/asset/", strconv.Itoa(realAssetIds[0]), gtu, gtp, http.StatusNotFound},
		{"User not same as auth", "/user/", strconv.Itoa(userIdOfAsset), "/asset/", strconv.Itoa(realAssetIds[0]), gtu, gtp, http.StatusUnauthorized},
		{"Existent User and Existent Asset but unauthorized", "/user/", strconv.Itoa(userIdOfAsset), "/asset/", strconv.Itoa(realAssetIds[0]), "admin", "123", http.StatusUnauthorized},
		{"Existent User and Existent Asset but unauthorized 2", "/user/", strconv.Itoa(userIdOfAsset), "/asset/", strconv.Itoa(realAssetIds[0]), "vas", "12345", http.StatusUnauthorized},
		{"Existent User and Existent Asset", "/user/", strconv.Itoa(userIdOfAsset), "/asset/", strconv.Itoa(realAssetIds[0]), gtuc, gtpc, http.StatusOK},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {

			wholeURL := srv.URL + tc.endPoint + tc.userid + tc.endPoint2 + tc.assetid + "/editdesc"

			type jsonExpected struct {
				NewDesc string `json:"newdesc"`
			}

			// so its always different
			body := &jsonExpected{
				NewDesc: "sth else ",
			}

			buf := new(bytes.Buffer)
			json.NewEncoder(buf).Encode(body)

			// Create client
			client := &http.Client{}

			// res, err := http.Delete(wholeURL)
			req, err := http.NewRequest("PUT", wholeURL, buf)
			if err != nil {
				t.Fatalf("could not send PUT request: %v", err)
			}
			req.SetBasicAuth(tc.username, tc.password)
			res, err := client.Do(req)
			if err != nil {
				t.Fatalf("Error: %v", err)
			}

			if res.StatusCode != tc.statusExpected {
				t.Errorf(" %v : Expected status %v got %v", wholeURL, tc.statusExpected, res.StatusCode)
			}

		})
	}
	db.removeAllAssets()
}
