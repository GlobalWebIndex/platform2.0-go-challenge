package models

type User struct {
	ID                uint                `gorm:"primary_key;"`
	Username          string              `gorm:"column:username;unique_index;not null;"`
	PasswordHash      string              `gorm:"column:password_hash;not null;"`
	FullName          string              `gorm:"column:full_name"`
	FavoriteCharts    []FavoritedChart    `gorm:"foreignkey:UserId"`
	FavoriteInsights  []FavoritedInsight  `gorm:"foreignkey:UserId"`
	FavoriteAudiences []FavoritedAudience `gorm:"foreignkey:UserId"`
}
