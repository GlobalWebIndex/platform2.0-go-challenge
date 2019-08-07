package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gwi/handler"
	"gwi/model"
	"gwi/utils"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	url = "http://localhost:1323/api/"
)

var users []model.User
var assets []model.Asset

func createUser() (*model.User, error) {

	url := url + "users/create"

	payload := strings.NewReader("{\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	body, _, err := execRequest(req)
	if err != nil {
		return nil, err
	}
	user := new(model.User)
	err = json.Unmarshal(body, user)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func getUser(id string) (*model.User, error) {

	url := url + "users/get/" + id
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	body, statusCode, err := execRequest(req)
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, utils.ErrNotFound
	}
	user := new(model.User)
	err = json.Unmarshal(body, user)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func getAssets(id string) (*model.Asset, error) {

	url := url + "assets/get/" + id

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	body, statusCode, err := execRequest(req)
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, utils.ErrNotFound
	}
	asset := new(model.Asset)
	err = json.Unmarshal(body, asset)
	if err != nil {
		return nil, err
	}
	return asset, nil

}

func createAsset(a model.Asset) (*model.Asset, error) {

	url := url + "assets/create"

	pl, err := json.Marshal(a)
	if err != nil {
		log.Fatal(err.Error())
	}
	payload := bytes.NewBuffer(pl)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	body, _, err := execRequest(req)
	if err != nil {
		return nil, err
	}
	asset := new(model.Asset)
	err = json.Unmarshal(body, asset)
	if err != nil {
		return nil, err
	}
	// fmt.Println(statusCode)
	// fmt.Println(string(body))
	return asset, nil

}

func updateAsset(a model.Asset) (*model.Asset, error) {

	url := url + "assets/"

	pl, err := json.Marshal(a)
	if err != nil {
		log.Fatal(err.Error())
	}
	payload := bytes.NewBuffer(pl)
	req, _ := http.NewRequest("PUT", url, payload)

	req.Header.Add("Content-Type", "application/json")
	body, _, err := execRequest(req)
	if err != nil {
		return nil, err
	}
	asset := new(model.Asset)
	err = json.Unmarshal(body, asset)
	if err != nil {
		return nil, err
	}
	// fmt.Println(statusCode)
	// fmt.Println(string(body))
	return asset, nil

}

func getUserFavourites(ufp handler.UserFavouritesPaged) (*model.User, error) {

	url := url + "users/favourites"

	pl, err := json.Marshal(ufp)
	if err != nil {
		log.Fatal(err.Error())
	}
	payload := bytes.NewBuffer(pl)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	body, statusCode, err := execRequest(req)
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, utils.ErrNotFound
	}
	u := new(handler.UserFavouritesPaged)
	err = json.Unmarshal(body, u)
	if err != nil {
		return nil, err
	}
	return &u.User, nil

}

func addUserFavorite(cdf handler.CDFavourite) error {

	url := url + "users/favourites"

	pl, err := json.Marshal(cdf)
	if err != nil {
		log.Fatal(err.Error())
	}
	payload := bytes.NewBuffer(pl)

	req, _ := http.NewRequest("PUT", url, payload)

	req.Header.Add("Content-Type", "application/json")
	_, statusCode, err := execRequest(req)
	if err != nil {
		return err
	}
	if statusCode != 200 {
		return utils.ErrAlreadyExists
	}
	// fmt.Println(statusCode)
	// fmt.Println(string(body))
	return nil

}

func removeUserFavorite(cdf handler.CDFavourite) error {

	url := url + "users/favourites"

	pl, err := json.Marshal(cdf)
	if err != nil {
		log.Fatal(err.Error())
	}
	payload := bytes.NewBuffer(pl)

	req, _ := http.NewRequest("DELETE", url, payload)

	req.Header.Add("Content-Type", "application/json")
	_, statusCode, err := execRequest(req)
	if err != nil {
		return err
	}
	if statusCode != 200 {
		return utils.ErrNotFound
	}
	return nil

}

func getAllUsers() ([]model.User, error) {
	users := make([]model.User, 0)
	url := url + "users"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	body, _, err := execRequest(req)
	if err != nil {
		return nil, err
	}
	// fmt.Println(statusCode)
	// fmt.Println(string(body))
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func execRequest(req *http.Request) ([]byte, int, error) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, 0, err
	}
	return body, res.StatusCode, nil
}

func userTests(loops int) []model.User {
	users, err := getAllUsers()
	if err != nil {
		panic(err)
	}
	startUsers := len(users)
	for i := 0; i < loops; i++ {
		user, err := createUser()
		if err != nil {
			panic(err)
		}
		users = append(users, *user)
	}
	users, err = getAllUsers()
	if err != nil {
		panic(err)
	}
	endUsers := len(users)
	if endUsers-startUsers != loops {
		panic("User Creation Error")
	}
	user, err := getUser(fmt.Sprint(users[endUsers-1].ID))
	if err != nil || user.ID != users[endUsers-1].ID {
		fmt.Println(err.Error())
		panic("User Recovery Error on Success")
	}
	user, err = getUser(fmt.Sprint(users[endUsers-1].ID + 50))
	if err == nil {
		panic("User Recovery Error on Not Found")
	}
	fmt.Println("User Tests Successful")
	return users
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}

func generateAsset() model.Asset {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	asset := model.Asset{Description: randStringRunes(10)}
	choice := r1.Int31n(3)
	switch choice {
	case 0:
		asset.Chart = &model.Chart{Title: randStringRunes(5), AxisTitles: [2]string{randStringRunes(5), randStringRunes(5)}}
	case 1:
		asset.Insight = &model.Insight{Insight: randStringRunes(5)}
	default:
		age := r1.Int31n(10)
		asset.Audience = &model.Audience{Characteristics: model.Characteristics{Gender: model.Other, BirthCountry: randStringRunes(2), AgeGroup: [2]uint{uint(age), uint(age + 10)}, HoursOnSM: uint(age), LastMoPurch: uint(age + 8)}}
	}
	return asset
}

func equalAsset(a1 model.Asset, a2 model.Asset) bool {
	var ins, char, aud bool
	if a1.Insight != nil && a2.Insight != nil {
		ins = *a1.Insight == *a2.Insight
	} else {
		ins = a1.Insight == a2.Insight
	}
	if a1.Audience != nil && a2.Audience != nil {
		aud = *a1.Audience == *a2.Audience
	} else {
		aud = a1.Audience == a2.Audience
	}
	if a1.Chart != nil && a2.Chart != nil {
		char = *a1.Chart == *a2.Chart
	} else {
		char = a1.Chart == a2.Chart
	}
	return a1.ID == a2.ID && a1.Description == a2.Description && ins && aud && char //&& *a1.Chart == *a2.Chart
}

func assetTests(loops int) {
	assets = make([]model.Asset, 0)
	for i := 0; i < loops; i++ {
		asset, err := createAsset(generateAsset())
		if err != nil {
			panic(err)
		}
		assets = append(assets, *asset)
	}
	for _, asset := range assets {
		as, err := getAssets(fmt.Sprint(asset.ID))
		if err != nil {
			panic(err)
		}
		if !equalAsset(*as, asset) {
			panic("Not the Same")
		}
	}
	assets[len(assets)-1].Description = "not a random desc"
	_, err := updateAsset(assets[len(assets)-1])
	if err != nil {
		panic(err)
	}
	as, err := getAssets(fmt.Sprint(assets[len(assets)-1].ID))
	if !equalAsset(*as, assets[len(assets)-1]) {
		panic("Not the Same")
	}
	_, err = getAssets(fmt.Sprint(assets[len(assets)-1].ID + 10))
	if err == nil {
		panic("Asset error on non existent")
	}
	fmt.Println("Asset Tests Successful")
}

func favoriteTests() {
	idx := len(users) - 1
	cdf := handler.CDFavourite{User: users[idx], Asset: assets[0]}
	err := addUserFavorite(cdf)
	if err != nil {
		panic("Favorite error on addition")
	}
	err = addUserFavorite(cdf)
	if err == nil {
		panic("Favorite error on existing addition")
	}
	ufp := handler.UserFavouritesPaged{User: users[idx], NextToken: 0, PageSize: 100}
	user, err := getUserFavourites(ufp)
	if err != nil || len(user.Assets) != 1 || !equalAsset(user.Assets[fmt.Sprint(assets[0].ID)], assets[0]) {
		panic("Favorite error on existing retrieve")
	}
	err = removeUserFavorite(cdf)
	if err != nil {
		panic("Favorite error on existing deletion")
	}
	err = removeUserFavorite(cdf)
	if err == nil {
		panic("Favorite error on non existing deletion")
	}
	fmt.Println("Favorite Tests Successful")
}

func main() {
	userLoops := 5
	assetLoops := 5
	users = userTests(userLoops)

	assetTests(assetLoops)
	favoriteTests()
}
