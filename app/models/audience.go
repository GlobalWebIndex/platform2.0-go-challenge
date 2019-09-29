package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Audience Section
type Audience struct {
	ID      uint32 `gorm:"primary_key"`
	AssetId uint32 `json:"assetId"`

	Gender      string `gorm:"size:1" json:"gender"`
	CountryCode string `gorm:"size:2" json:"country_code"`
	AgeFrom     uint8  `gorm:"size:2" json:"age_from"`
	AgeTo       uint8  `gorm:"size:2" json:"age_to"`
}

type AudienceDetails struct {
	Title  string
	Result float64
}

func (a *Audience) TableName() string {
	return "audiences"
}

func (a *Audience) CreateAssetAudience(db *gorm.DB, uid uint32) (*Asset, *Audience, error) {
	asset := &Asset{}
	asset.UserID = uid
	asset.AssetType = "audience"
	asset, err := asset.SaveAsset(db, uid)
	if err != nil {
		return nil, nil, err
	}

	a.AssetId = asset.ID
	err = db.Create(&a).Error
	if err != nil {
		return asset, a, err
	}
	return asset, a, nil
}

func GetAssetAudienceGopher(db *gorm.DB, id uint32, a Asset, cAsset chan Asset) {
	audience := Audience{}
	db.Model(&Audience{}).Where("asset_id = ?", id).Take(&audience)
	a.AudienceData = &audience

	audienceDetails := AudienceDetails{}

	type SqlResult struct {
		Avg float64
	}
	var result SqlResult
	db.Raw(`SELECT AVG(hours_spent_on_social_daily) from participants 
		WHERE
		(age >= ? OR ? = 0)
		AND (age <= ? OR ? = 0)
		AND (country_code >= ? OR ? = '')
		AND (gender = ? OR ? = '')
		`, audience.AgeFrom, audience.AgeFrom,
		audience.AgeTo, audience.AgeTo,
		audience.CountryCode, audience.CountryCode,
		audience.Gender, audience.Gender).Scan(&result)
	fmt.Println("Result for audience", result)
	audienceDetails.Result = result.Avg

	a.Audience = &AudienceDetails{
		Title:  AudienceTitle(audience, audienceDetails),
		Result: result.Avg,
	}

	cAsset <- a
}
