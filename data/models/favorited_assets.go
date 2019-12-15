package models

type FavoritedChart struct {
	ChartID              uint   `gorm:"column:chart_id;primary_key;auto_increment:false;"`
	UserID               uint   `gorm:"column:user_id;primary_key;auto_increment:false;"`
	FavoritedDescription string `gorm:"column:favorited_description;"`
}

type FavoritedInsight struct {
	InsightID            uint   `gorm:"column:insight_id;primary_key;auto_increment:false;"`
	UserID               uint   `gorm:"column:user_id;primary_key;auto_increment:false;"`
	FavoritedDescription string `gorm:"column:favorited_description;"`
}

type FavoritedAudience struct {
	AudienceID           uint   `gorm:"column:audience_id;primary_key;auto_increment:false;"`
	UserID               uint   `gorm:"column:user_id;primary_key;auto_increment:false;"`
	FavoritedDescription string `gorm:"column:favorited_description;"`
}
