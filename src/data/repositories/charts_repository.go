package repositories

import (
	"gwi-challenge/data"
	"gwi-challenge/data/models"
)

// GetChartByTitle returns user with specified title
func GetChartByTitle(title string) (chart *models.Chart, err error) {
	chart = new(models.Chart)
	db := data.GetDB()
	err = db.Preload("Data").First(chart, &models.Chart{Title: title}).Error
	return
}

// DeleteChartPoint does what its name says
func DeleteChartPoint(point *models.ChartPoint) error {
	db := data.GetDB()
	err := db.Delete(point).Error
	return err
}

func GetFavoritedChart(userID uint, chartID uint) (favoritedChart *models.FavoritedChart, err error) {
	favoritedChart = &models.FavoritedChart{}
	db := data.GetDB()
	err = db.First(favoritedChart, &models.FavoritedChart{UserID: userID, ChartID: chartID}).Error
	return
}

func GetChartsByIds(ids []uint, charts *[]models.Chart) (err error) {
	*charts = []models.Chart{}
	db := data.GetDB()
	err = db.Preload("Data").Where(ids).Find(charts).Error
	return
}

func GetChartsByIdsPaginated(ids []uint, charts *[]models.Chart, pageSize uint, offset uint) (err error) {
	*charts = []models.Chart{}
	db := data.GetDB()
	err = db.Preload("Data").Where(ids).Limit(pageSize).Offset(offset).Find(charts).Error
	return
}

func GetAllCharts(charts *[]models.Chart) (err error) {
	*charts = []models.Chart{}
	db := data.GetDB()
	err = db.Preload("Data").Find(charts).Error
	return
}

func GetAllChartsPaginated(charts *[]models.Chart, pageSize uint, offset uint) (err error) {
	*charts = []models.Chart{}
	db := data.GetDB()
	err = db.Preload("Data").Limit(pageSize).Offset(offset).Find(charts).Error
	return
}
