package repositories

import (
	"gwi-challenge/data"
	"gwi-challenge/data/models"
)

func GetAudienceByTitle(title string) (audience *models.Audience, err error) {
	audience = new(models.Audience)
	db := data.GetDB()
	err = db.Preload("AudienceInfo").First(audience, &models.Audience{Title: title}).Error
	return
}

func DeleteAudience(audience *models.Audience) error {
	db := data.GetDB()
	err := db.Delete(audience).Error
	return err
}

func GetFavoritedAudience(userID uint, audienceID uint) (favoritedAudience *models.FavoritedAudience, err error) {
	favoritedAudience = new(models.FavoritedAudience)
	db := data.GetDB()
	err = db.First(favoritedAudience, &models.FavoritedAudience{UserID: userID, AudienceID: audienceID}).Error
	return
}

func GetAudiencesByIds(ids []uint, audiences *[]models.Audience) (err error) {
	*audiences = []models.Audience{}
	db := data.GetDB()
	err = db.Preload("AudienceInfo").Where(ids).Find(audiences).Error
	return
}

func GetAudiencesByIdsPaginated(ids []uint, audiences *[]models.Audience, pageSize uint, offset uint) (err error) {
	*audiences = []models.Audience{}
	db := data.GetDB()
	err = db.Preload("AudienceInfo").Where(ids).Limit(pageSize).Offset(offset).Find(audiences).Error
	return
}

func GetAllAudiences(audiences *[]models.Audience) (err error) {
	*audiences = []models.Audience{}
	db := data.GetDB()
	err = db.Preload("AudienceInfo").Find(audiences).Error
	return
}

func GetAllAudiencesPaginated(audiences *[]models.Audience, pageSize uint, offset uint) (err error) {
	*audiences = []models.Audience{}
	db := data.GetDB()
	err = db.Preload("AudienceInfo").Limit(pageSize).Offset(offset).Find(audiences).Error
	return
}
