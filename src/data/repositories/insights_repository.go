package repositories

import (
	"gwi-challenge/data"
	"gwi-challenge/data/models"
)

func GetInsightByTitle(title string) (insight *models.Insight, err error) {
	insight = new(models.Insight)
	db := data.GetDB()
	err = db.First(insight, &models.Insight{Title: title}).Error
	return
}

func GetFavoritedInsight(userID uint, isnightID uint) (favoritedInsight *models.FavoritedInsight, err error) {
	favoritedInsight = new(models.FavoritedInsight)
	db := data.GetDB()
	err = db.First(favoritedInsight, &models.FavoritedInsight{UserID: userID, InsightID: isnightID}).Error
	return
}

func GetInsightsByIds(ids []uint, insights *[]models.Insight) (err error) {
	*insights = []models.Insight{}
	db := data.GetDB()
	err = db.Where(ids).Find(insights).Error
	return
}

func GetInsightsByIdsPaginated(ids []uint, insights *[]models.Insight, pageSize uint, offset uint) (err error) {
	*insights = []models.Insight{}
	db := data.GetDB()
	err = db.Where(ids).Limit(pageSize).Offset(offset).Find(insights).Error
	return
}

func GetAllInsights(insights *[]models.Insight) (err error) {
	*insights = []models.Insight{}
	db := data.GetDB()
	err = db.Find(insights).Error
	return
}

func GetAllInsightsPaginated(insights *[]models.Insight, pageSize uint, offset uint) (err error) {
	*insights = []models.Insight{}
	db := data.GetDB()
	err = db.Limit(pageSize).Offset(offset).Find(insights).Error
	return
}
