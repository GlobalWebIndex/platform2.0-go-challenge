package main

import (
	"math/rand"
	"time"
)

// a seed randomiser based on time
var r = rand.New(rand.NewSource(1)) //keep this to perform the same randoms
// var r = rand.New(rand.NewSource(time.Now().UnixNano())) // true random random

// counter to keep number of assets
var assetCounter = 0

// create a random string of n length
func randomString(n int) string {
	// in case a letter should appear more frequently than others just add it multiple times to the pile of letters
	var pile = []byte(" abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, n)
	for i := range b {
		b[i] = pile[r.Intn(len(pile))]
	}
	return string(b)
}

// create a random point
func getRandPoint() Point {
	return Point{
		X: 10 * r.Float64(),
		Y: 100 * r.Float64(),
	}
}

// create a slice of random points
func getRandPointList(n int) []Point {
	randPoints := []Point{}
	for i := 0; i < n; i++ {
		randPoints = append(randPoints, getRandPoint())
	}
	return randPoints
}

// create a slice of random strings
func genRandTags(n int, tagSize int) []string {
	randTags := []string{}
	for i := 0; i < n; i++ {
		randTags = append(randTags, randomString(tagSize))
	}
	return randTags
}

// create a random struct of Atributes
func genRandAttributes(aType AssetType) *CommonAttributes {
	// increment number of assets
	assetCounter++
	return &CommonAttributes{
		ID:      assetCounter,
		UserID:  r.Intn(50) + 1, // user ids from [1-50]
		Desc:    randomString(r.Intn(100)),
		Type:    aType,
		CrDate:  time.Now(),
		LastUpd: time.Now(),
		IsFav:   r.Float32() < 0.5,
		Stars:   r.Intn(5),
	}
}

//create a random struct of Chart
func genRandChart() Chart {
	return Chart{
		Attr:  genRandAttributes(CHART),
		Title: randomString(40),
		XAxis: randomString(10),
		YAxis: randomString(10),
		Data:  getRandPointList(r.Intn(50)),
	}
}

//create a random struct of Insight
func genRandInsight() Insight {
	return Insight{
		Attr:   genRandAttributes(INSIGHT),
		InsMsg: randomString(40),
		Tags:   genRandTags(r.Intn(5), 15),
	}
}

//create a random struct of Audience
func genRandAudience() Audience {
	randomGender := MALE
	if r.Float32() < 0.5 {
		randomGender = FEMALE
	}
	ageStart := r.Intn(5)
	return Audience{
		Attr:            genRandAttributes(AUDIENCE),
		Gen:             randomGender,
		BrthCntry:       randomString(10),
		AgeGrpFrom:      ageStart,
		AgeGrpTo:        ageStart + r.Intn(15),
		HSpentSM:        10 * r.Float32(),
		NumPurchsLMonth: r.Intn(5),
	}
}

//create a random Asset
func genRandAsset() Asset {
	typeAsset := r.Intn(3)

	switch typeAsset {
	case 0:
		return genRandChart()
	case 1:
		return genRandInsight()
	case 2:
		return genRandAudience()
	}

	return genRandChart()

}

// a function to create n assets
func fillAssets(n int) {

	for i := 0; i < n; i++ {
		// create the asset
		asset := genRandAsset()
		// add it
		DBaddAsset(&asset)
	}
}
