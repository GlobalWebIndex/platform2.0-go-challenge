package services

import (
	"errors"
	"gwi-challenge/data/models"
	"gwi-challenge/data/repositories"
	"strings"
)

func GetChartByTitle(title string) (chartToReturn *models.Chart, err error) {
	chartToReturn, err = repositories.GetChartByTitle(title)
	return
}

func CreateNewChart(chartToCreate *models.Chart) error {
	err := repositories.Save(&chartToCreate)
	if err != nil && strings.Contains(err.Error(), "constraint") {
		err = errors.New("title already exists")
	}
	return err
}

func UpdateChartByTitle(title string, newChart *models.Chart) (updatedChart *models.Chart, err error) {
	updatedChart, err = repositories.GetChartByTitle(title)
	if err != nil {
		return
	}
	err = deleteAllChartPoints(updatedChart)
	if err != nil {
		return
	}
	updatedChart.Title = newChart.Title
	updatedChart.XAxes = newChart.XAxes
	updatedChart.YAxes = newChart.YAxes
	updatedChart.Data = newChart.Data
	err = repositories.Save(&updatedChart)
	return
}

func DeleteChartByTitle(title string) (err error) {
	chartToDelete, err := repositories.GetChartByTitle(title)
	if err != nil {
		return
	}
	err = deleteAllChartPoints(chartToDelete)
	if err != nil {
		return
	}
	err = repositories.Delete(&chartToDelete)
	return
}

func FavoriteChart(userID uint, title string) (err error) {
	chartToFavorite, err := GetChartByTitle(title)
	if err != nil {
		err = errors.New("Chart with title " + title + " does not exist")
		return
	}
	userWhoLoves, err := GetUserById(userID, false)
	if err != nil {
		err = errors.New("User does not exist")
		return
	}
	favoriteModel := models.FavoritedChart{
		ChartID: chartToFavorite.ID,
		UserID:  userWhoLoves.ID,
	}
	err = repositories.Save(&favoriteModel)
	return
}

func UnfavoriteChart(userID uint, title string) (err error) {
	chartToUnfavorite, err := GetChartByTitle(title)
	if err != nil {
		err = errors.New("Chart with title " + title + " does not exist")
		return
	}
	userWhoHates, err := GetUserById(userID, false)
	if err != nil {
		err = errors.New("User does not exist")
		return
	}
	favoritedModel, err := repositories.GetFavoritedChart(userWhoHates.ID, chartToUnfavorite.ID)
	if err != nil {
		err = errors.New("User has not favorited this chart")
		return
	}
	err = repositories.Delete(&favoritedModel)
	return
}

func DescribeChart(userID uint, title string, description string) (err error) {
	chartToPatch, err := GetChartByTitle(title)
	if err != nil {
		err = errors.New("Chart with title " + title + " does not exist")
		return
	}
	user, err := GetUserById(userID, false)
	if err != nil {
		err = errors.New("User does not exist")
		return
	}
	favoritedModel, err := repositories.GetFavoritedChart(user.ID, chartToPatch.ID)
	if err != nil {
		err = errors.New("User has not favorited this chart")
		return
	}
	favoritedModel.FavoritedDescription = description
	err = repositories.Save(&favoritedModel)
	return
}

func deleteAllChartPoints(chart *models.Chart) (err error) {
	errors := make(chan error)
	for _, chartPoint := range chart.Data {
		go deleteChartPoint(chartPoint, errors)
	}
	for i := 0; i < len(chart.Data); i++ {
		routineError := <-errors
		if routineError != nil {
			err = routineError
		}
	}
	close(errors)
	return err
}

func deleteChartPoint(point models.ChartPoint, c chan error) {
	err := repositories.DeleteChartPoint(&point)
	c <- err
}

func getAllCharts(charts *[]models.Chart, errs chan error) {
	err := repositories.GetAllCharts(charts)
	errs <- err
}

func getAllChartsPaginated(charts *[]models.Chart, pageSize uint, offset uint, errs chan error) {
	err := repositories.GetAllChartsPaginated(charts, pageSize, offset)
	errs <- err
}

func getChartsByIds(favoritedCharts *[]models.FavoritedChart, charts *[]models.Chart, errs chan error) {
	chartIds := []uint{}
	for _, chart := range *favoritedCharts {
		chartIds = append(chartIds, chart.ChartID)
	}
	err := repositories.GetChartsByIds(chartIds, charts)
	errs <- err
}

func getChartsByIdsPaginated(favoritedCharts *[]models.FavoritedChart, charts *[]models.Chart, pageSize uint, offset uint, errs chan error) {
	chartIds := []uint{}
	for _, chart := range *favoritedCharts {
		chartIds = append(chartIds, chart.ChartID)
	}
	err := repositories.GetChartsByIdsPaginated(chartIds, charts, pageSize, offset)
	errs <- err
}