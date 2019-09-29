package models

import (
	"github.com/jinzhu/gorm"
)

// Chart Section
type Chart struct {
	ID            uint32 `gorm:"primary_key"`
	AssetId       uint32 `json:"assetId"`
	GroupedMetric string `gorm:"size:30" json:"groupedMetric"` // Could be age,
}

type ChartDetails struct {
	XTitle string
	YTitle string
	Data   []ChartPoint
}
type ChartPoint struct {
	XValue string
	YValue float64
}

func (c *Chart) TableName() string {
	return "charts"
}

func (c *Chart) CreateAssetChart(db *gorm.DB, uid uint32) (*Asset, *Chart, error) {
	asset := &Asset{}
	asset.UserID = uid
	asset.AssetType = "chart"
	asset, err := asset.SaveAsset(db, uid)
	if err != nil {
		return nil, nil, err
	}

	c.AssetId = asset.ID
	err = db.Create(&c).Error
	if err != nil {
		return asset, c, err
	}
	return asset, c, nil
}

func GetAssetChartGopher(db *gorm.DB, id uint32, a Asset, cAsset chan Asset) {
	chart := Chart{}
	db.Model(&Chart{}).Where("asset_id = ?", id).Take(&chart)
	a.ChartData = &chart

	chartDetails := ChartDetails{}

	var grouped string
	var avg float64
	rows, _ := db.Raw("select age, AVG(hours_spent_on_social_daily) from participants group by age order by age").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&grouped, &avg)
		chartDetails.Data = append(chartDetails.Data, ChartPoint{XValue: grouped, YValue: avg})
	}
	chartDetails.XTitle = chart.GroupedMetric
	chartDetails.YTitle = "age"

	a.Chart = &chartDetails
	cAsset <- a
}
