package models

import (
	"github.com/jinzhu/gorm"
)

// Insight Section
type Insight struct {
	ID      uint32 `gorm:"primary_key"`
	AssetId uint32 `json:"assetId"`

	Gender      string `gorm:"size:1" json:"gender"`
	CountryCode string `gorm:"size:2" json:"country_code"`
	AgeFrom     uint8  `gorm:"size:2" json:"age_from"`
	AgeTo       uint8  `gorm:"size:2" json:"age_to"`

	HoursComparator string `gorm:"size:2" json:"hours_comparator"`
	HoursReference  uint8  `gorm:"size:1" json:"hours_reference_point"`
}

type InsightDetails struct {
	Title          string
	Sample         uint64
	FiltererSample uint64
}

func (i *Insight) TableName() string {
	return "insights"
}

func (i *Insight) CreateAssetInsight(db *gorm.DB, uid uint32) (*Asset, *Insight, error) {
	asset := &Asset{}
	asset.UserID = uid
	asset.AssetType = "insight"
	asset, err := asset.SaveAsset(db, uid)
	if err != nil {
		return nil, nil, err
	}

	i.AssetId = asset.ID
	err = db.Create(&i).Error
	if err != nil {
		return asset, i, err
	}

	return asset, i, nil
}

func GetAssetInsightGopher(db *gorm.DB, id uint32, a Asset, cAsset chan Asset) {
	insight := Insight{}
	db.Debug().Model(&Insight{}).Where("asset_id = ?", id).Take(&insight)
	a.InsightData = &insight

	insightDetails := InsightDetails{}

	type SqlResult struct {
		Count uint64
	}
	var sample, sampleF SqlResult
	db.Raw(`select COUNT(*) from participants 
		WHERE
		(age >= ? OR ? = 0)
		AND (age >= ? OR ? = 0)
		AND (country_code = ? OR ? = '')
		AND (gender = ? OR ? = '')
		`, insight.AgeFrom, insight.AgeFrom,
		insight.AgeTo, insight.AgeTo,
		insight.CountryCode, insight.CountryCode,
		insight.Gender, insight.Gender).
		Scan(&sample)

	db.Raw(`select COUNT(*) from participants 
		WHERE
		(age >= ? OR ? = 0)
		AND (age <= ? OR ? = 0)
		AND (country_code = ? OR ? = '')
		AND (gender = ? OR ? = '')
		AND (hours_spent_on_social_daily > ? OR ? != '>')
		AND (hours_spent_on_social_daily < ? OR ? != '<')
		AND (hours_spent_on_social_daily >= ? OR ? != '>=')
		AND (hours_spent_on_social_daily <= ? OR ? != '<=')
		`, insight.AgeFrom, insight.AgeFrom,
		insight.AgeTo, insight.AgeTo,
		insight.CountryCode, insight.CountryCode,
		insight.Gender, insight.Gender,
		insight.HoursReference, insight.HoursComparator,
		insight.HoursReference, insight.HoursComparator,
		insight.HoursReference, insight.HoursComparator,
		insight.HoursReference, insight.HoursComparator).
		Scan(&sampleF)

	insightDetails.Sample = sample.Count
	insightDetails.FiltererSample = sampleF.Count

	a.Insight = &InsightDetails{
		Title:          InsightTitle(insight, insightDetails),
		Sample:         sample.Count,
		FiltererSample: sampleF.Count,
	}

	cAsset <- a
}
