package models

type Insight struct {
	ID          uint               `gorm:"primary_key;"`
	FavoritedBy []FavoritedInsight `gorm:"foreignkey:InsightID"`
	Title       string             `gorm:"column:title;unique_index;not null;"`
	Text        string             `gorm:"column:insight_text;"`
}
