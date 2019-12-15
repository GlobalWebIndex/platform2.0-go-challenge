package models

type Chart struct {
	ID          uint             `gorm:"primary_key;"`
	FavoritedBy []FavoritedChart `gorm:"foreignkey:ChartID"`
	Title       string           `gorm:"column:title;unique_index;not null;"`
	XAxes       string
	YAxes       string
	Data        []ChartPoint `gorm:"foreignkey:ChartID"`
}

type ChartPoint struct {
	ID      uint `gorm:"primary_key;"`
	ChartID uint `gorm:"column:chart_id"`
	X       float32
	Y       float32
}
