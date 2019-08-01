package main

type assetType int32

const (
	typeChart assetType = iota
	typeInsight
	typeAudience
)

type asset struct {
	ID          string `json:"id"`
	UserID      string `json:"userId"`
	AssetType   int32  `json:"assetType"`
	Favorite    bool   `json:"favorite"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Data        string `json:"data"`
}

type chart struct {
	Title  string  `json:"title"`
	XTitle string  `json:"xTitle"`
	YTitle string  `json:"yTitle"`
	Points []point `json:"points"`
}
type point struct {
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
	Value       float64 `json:"value"`
	Description float64 `json:"description"`
}

type insight struct {
	Value string `json:"value"`
}

type gender int32

const (
	male gender = iota
	female
)

type audience struct {
	Gender            gender `json:"gender"`
	BirthCountry      string `json:"birthCountry"`
	AgeGroupMin       int32  `json:"ageGroupMin"`
	AgeGroupMax       int32  `json:"ageGroupMax"`
	HoursOnSocial     int32  `json:"hoursOnSocial"`
	NumberOfPurchases int32  `json:"numberOfPurchases"`
}

type editAssetRequest struct {
	Description string `json:"description"`
}

var assets = []asset{
	asset{UserID: "1", ID: "1", Title: "Chart 1", Description: "Chart 1 Description", AssetType: 1, Favorite: true, Data: `{
"title": "Chart title",
"xTitle": "X Title",
"yTitle": "Y Title",
"points": [
{"x":0,"y":1,"value":45,"description":"Point Description"},
{"x":1,"y":2,"value":55,"description":"Point 2 Description"}
]
}`},
	asset{UserID: "1", ID: "2", Title: "Insight 1", Description: "Insight 1 Description", AssetType: 2, Favorite: false, Data: `{"value":"Insight 1 description is great !"}`},
	asset{UserID: "1", ID: "3", Title: "Audience 1", Description: "Audience 1 Description", AssetType: 3, Favorite: true, Data: `{
"gender": 1
"birthCountry":"gr"
"ageGroupMin":20
"ageGroupMax":25
"hoursOnSocial":0.5
"numberOfPurchases":1
}`},
}
