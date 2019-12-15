package repositories

import (
	"gwi-challenge/data"
	"gwi-challenge/data/models"
)

// GetUserByUsername returns user with specified username
func GetUserByUsername(username string, withFavorites bool) (user *models.User, err error) {
	user = &models.User{}
	db := data.GetDB()
	err = db.First(user, &models.User{Username: username}).Error
	return
}

// GetUserByID returns user with specified id
func GetUserByID(id uint, withFavorites bool) (user *models.User, err error) {
	user = &models.User{}
	db := data.GetDB()
	if withFavorites {
		err = db.Preload("FavoriteCharts").Preload("FavoriteInsights").Preload("FavoriteAudiences").First(user, &models.User{ID: id}).Error
	} else {
		err = db.First(user, &models.User{ID: id}).Error
	}
	return
}
