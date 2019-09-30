package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Asset struct {
	ID             uint32 `gorm:"primary_key;auto_increment" json:"id"`
	UserID         uint32 `json:"uid"`
	IsFavorite     bool   `gorm:"not null" json:"is_favorite"`
	Title          string `gorm:"size:255;not null" json:"title"`
	TitleGenerated string `gorm:"-" json:"title_generated"`
	AssetType      string `gorm:"size:15;not null" json:"asset_type"`

	ChartData    *Chart    `json:",omitempty"`
	InsightData  *Insight  `json:",omitempty"`
	AudienceData *Audience `json:",omitempty"`

	Chart    *ChartDetails    `json:",omitempty"`
	Insight  *InsightDetails  `json:",omitempty"`
	Audience *AudienceDetails `json:",omitempty"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type AssetChannelResult struct {
	Result Asset
	Err    error
}

func (a *Asset) TableName() string {
	return "assets"
}

func (a *Asset) Prepare() {
	a.ID = 0
	// a.Title = html.EscapeString(strings.TrimSpace(a.Title))
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
}

func (a *Asset) GetAssets(db *gorm.DB, uid uint32, page, itemsPerPage uint32) ([]Asset, error) {
	assets := []Asset{}
	err := db.Debug().Model(&Asset{}).Order("updated_at DESC").Offset((page-1)*itemsPerPage).Limit(itemsPerPage).Where("user_id = ?", uid).Find(&assets).Error
	if err != nil {
		return []Asset{}, err
	}
	return assets, nil
}

func (a *Asset) GetAssetsWithDetails(db *gorm.DB, assets []Asset, uid uint32) ([]Asset, error) {
	assetsResponse := []Asset{}
	var err error

	// Initial Implementation to check Data returned ~200+ms
	// for i := 0; i < len(ids); i++ {
	// 	asset, _ := GetAsset(db, ids[i], uid)
	// 	assetsResponse = append(assetsResponse, asset)
	// }

	// Initial Implementation 250ms - Bellow Implementation ~30-60ms
	channelAsset := make(chan Asset)
	for _, a := range assets {
		switch a.AssetType {
		case "chart":
			go GetAssetChartGopher(db, a.ID, a, channelAsset)
		case "insight":
			go GetAssetInsightGopher(db, a.ID, a, channelAsset)
		case "audience":
			go GetAssetAudienceGopher(db, a.ID, a, channelAsset)
		default:
		}
	}
	for i := 0; i < len(assets); i++ {
		assetsResponse = append(assetsResponse, <-channelAsset)
	}

	return assetsResponse, err
}

func (a *Asset) UpdateAsset(db *gorm.DB, id uint32, uid uint32) (*Asset, error) {

	db = db.Debug().Model(&User{}).Where("id = ? and user_id = ?", id, uid).Take(&Asset{}).UpdateColumns(
		map[string]interface{}{
			"is_favorite": a.IsFavorite,
			"title":       a.Title,
			"update_at":   time.Now(),
		},
	)
	if db.Error != nil {
		return &Asset{}, db.Error
	}
	// This is the display the updated user
	err := db.Debug().Model(&Asset{}).Where("id = ?", id).Take(&a).Error
	if err != nil {
		return &Asset{}, err
	}
	return a, nil
}

func (a *Asset) SaveAsset(db *gorm.DB, uid uint32) (*Asset, error) {
	err := db.Model(&Asset{}).Create(&a).Error
	if err != nil {
		return &Asset{}, err
	}
	if a.ID != 0 {
		err = db.Model(&Asset{}).Where("id = ?", a.ID).Error
		if err != nil {
			return &Asset{}, err
		}
	}
	return a, nil
}

func GetAsset(db *gorm.DB, id uint32, uid uint32) (Asset, error) {
	a := Asset{}
	err := db.Debug().Model(&Asset{}).Where("id = ? and user_id = ?", id, uid).Take(&a).Error
	if err != nil {
		return Asset{}, err
	}
	if a.ID != 0 {
		err = db.Debug().Model(&Asset{}).Where("id = ? and user_id = ?", a.Title).Error
		if err != nil {
			return Asset{}, err
		}
	}

	var errRes error
	channelAsset := make(chan Asset)
	switch a.AssetType {
	case "chart":
		go GetAssetChartGopher(db, a.ID, a, channelAsset)
	case "insight":
		go GetAssetInsightGopher(db, a.ID, a, channelAsset)
	case "audience":
		go GetAssetAudienceGopher(db, a.ID, a, channelAsset)
	default:
	}
	a = <-channelAsset
	if errRes != nil {
		return Asset{}, errRes
	}
	return a, errRes
}

func (a *Asset) DeleteAsset(db *gorm.DB, id uint32, uid uint32) (int64, error) {
	var rowsAffected int64
	db = db.Debug().Model(&Asset{}).Where("id = ? and user_id = ?", id, uid).Take(&Asset{}).Delete(&Asset{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Asset not found")
		}
		return 0, db.Error
	}
	rowsAffected = db.RowsAffected

	if a.AssetType == "chart" {
		db.Debug().Model(&Chart{}).Where("asset_id = ?").Take(&Chart{}).Delete(&Asset{})
	}
	if a.AssetType == "insight" {
		db.Debug().Model(&Insight{}).Where("asset_id = ?").Take(&Insight{}).Delete(&Asset{})
	}
	if a.AssetType == "audience" {
		db.Debug().Model(&Audience{}).Where("asset_id = ?").Take(&Audience{}).Delete(&Asset{})
	}

	return rowsAffected, nil
}
