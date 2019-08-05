package main

import (
	"time"
)

type Asset interface {
	getAttributes() *CommonAttributes
}

func getAssetUserId(a *Asset) int {
	return (*a).getAttributes().UserID
}

func getAssetId(a *Asset) int {
	return (*a).getAttributes().ID
}

func getAssetDesc(a *Asset) string {
	return (*a).getAttributes().Desc
}

func getUserId(a *Asset) int {
	return (*a).getAttributes().UserID
}

func (db *dbMock) EditDesc(a *Asset, newDesc string) {
	(*a).getAttributes().Desc = newDesc
	(*a).getAttributes().LastUpd = time.Now()
	db.DBupdateAssetPersist(a)
}

func (db *dbMock) FavorAsset(a *Asset) {
	(*a).getAttributes().IsFav = true
	(*a).getAttributes().LastUpd = time.Now()
	db.DBupdateAssetPersist(a)
}

func (db *dbMock) UnFavorAsset(a *Asset) {
	(*a).getAttributes().IsFav = false
	(*a).getAttributes().LastUpd = time.Now()
	db.DBupdateAssetPersist(a)
}

// User the owner of the asset
type User struct {
	UserID   int    `json:"userid"`
	Username string `json:"username"`
	Password string `json:"password"` // it should be md5
}

// CommonAttributes : keep common attributes of assets in a seperate struct
type CommonAttributes struct {
	ID      int       `json:"assetid"`
	UserID  int       `json:"userid"`
	Desc    string    `json:"desc"`
	Type    AssetType `json:"assettype"`
	CrDate  time.Time `json:"crdate"`
	LastUpd time.Time `json:"lastupd"`
	IsFav   bool      `json:"isfav"`
	Stars   int       `json:"stars"`
}

// Point : simple 2D data point
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Chart : definition of Chart
type Chart struct {
	Attr  *CommonAttributes `json:"attr"`
	Title string            `json:"title"`
	XAxis string            `json:"xaxis"`
	YAxis string            `json:"yaxis"`
	Data  []Point           `json:"data"`
}

func (c Chart) getType() string {
	return "Chart"
}

func (c Chart) getAttributes() *CommonAttributes {
	return c.Attr
}

// Insight : definition of Insight
type Insight struct {
	Attr   *CommonAttributes `json:"attr"`
	InsMsg string            `json:"insightmessage"`
	Tags   []string          `json:"tags"`
}

func (i Insight) getAttributes() *CommonAttributes {
	return i.Attr
}

type Gender string

const (
	MALE   Gender = "MALE"
	FEMALE Gender = "FEMALE"
)

type AssetType string

const (
	CHART    AssetType = "CHART"
	INSIGHT  AssetType = "INSIGHT"
	AUDIENCE AssetType = "AUDIENCE"
)

// Audience : definition of Audience
type Audience struct {
	Attr            *CommonAttributes `json:"attr"`
	Gen             Gender            `json:"gender"`
	BrthCntry       string            `json:"birthCountry"`
	AgeGrpFrom      int               `json:"agegroupfrom"`
	AgeGrpTo        int               `json:"agegroupto"`
	HSpentSM        float32           `json:"hoursSpendSocialMedia"`
	NumPurchsLMonth int               `json:"NumbersPurchasesLMonth"`
}

func (a Audience) getType() string {
	return "Audience"
}

func (a Audience) getAttributes() *CommonAttributes {
	return a.Attr
}

type AssetGroupped struct {
	CH []*Asset `json:"chartsl"`
	IN []*Asset `json:"insighsl"`
	AU []*Asset `json:"audiencesl"`
}

// a useful function that allows to group assets
// by their assettype
func getAssetsGrouped(allAssets map[int]*Asset) AssetGroupped {

	aGrp := AssetGroupped{}
	for _, a := range allAssets {
		switch (*a).getAttributes().Type {
		case CHART:
			aGrp.CH = append(aGrp.CH, a)
		case INSIGHT:
			aGrp.IN = append(aGrp.IN, a)
		case AUDIENCE:
			aGrp.AU = append(aGrp.AU, a)
		}
	}
	return aGrp
}
