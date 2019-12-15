package services

import (
	"errors"
	"gwi-challenge/data/models"
	"gwi-challenge/data/repositories"
	"strings"
)

func GetInsightByTitle(title string) (insightToReturn *models.Insight, err error) {
	insightToReturn, err = repositories.GetInsightByTitle(title)
	return
}

func CreateNewInsight(insightToCreate *models.Insight) error {
	err := repositories.Save(&insightToCreate)
	if err != nil && strings.Contains(err.Error(), "constraint") {
		err = errors.New("title already exists")
	}
	return err
}

func UpdateInsightByTitle(title string, newInsight *models.Insight) (updatedInsight *models.Insight, err error) {
	updatedInsight, err = repositories.GetInsightByTitle(title)
	if err != nil {
		return
	}
	updatedInsight.Title = newInsight.Title
	updatedInsight.Text = newInsight.Text
	err = repositories.Save(&updatedInsight)
	if err != nil && strings.Contains(err.Error(), "constraint") {
		err = errors.New("title already exists")
	}
	return
}

func DeleteInsightByTitle(title string) (err error) {
	insightToDelete, err := repositories.GetInsightByTitle(title)
	if err != nil {
		return
	}
	err = repositories.Delete(&insightToDelete)
	return
}

func FavoriteInsight(userID uint, title string) (err error) {
	insightToFavorite, err := GetInsightByTitle(title)
	if err != nil {
		err = errors.New("Insight with title " + title + " does not exist")
		return
	}
	userWhoLoves, err := GetUserById(userID, false)
	if err != nil {
		err = errors.New("User does not exist")
		return
	}
	favoriteModel := models.FavoritedInsight{
		InsightID: insightToFavorite.ID,
		UserID:    userWhoLoves.ID,
	}

	err = repositories.Save(&favoriteModel)
	return
}

func UnfavoriteInsight(userID uint, title string) (err error) {
	insightToUnfavorite, err := GetInsightByTitle(title)
	if err != nil {
		err = errors.New("Insight with title " + title + " does not exist")
		return
	}
	userWhoHates, err := GetUserById(userID, false)
	if err != nil {
		err = errors.New("User does not exist")
		return
	}

	favoritedModel, err := repositories.GetFavoritedInsight(userWhoHates.ID, insightToUnfavorite.ID)
	if err != nil {
		err = errors.New("User has not favorited this insight")
		return
	}
	err = repositories.Delete(&favoritedModel)
	return
}

func DescribeInsight(userID uint, title string, description string) (err error) {
	insightToPatch, err := GetInsightByTitle(title)
	if err != nil {
		err = errors.New("Insight with title " + title + " does not exist")
		return
	}
	user, err := GetUserById(userID, false)
	if err != nil {
		err = errors.New("User does not exist")
		return
	}

	favoritedModel, err := repositories.GetFavoritedInsight(user.ID, insightToPatch.ID)
	if err != nil {
		err = errors.New("User has not favorited this insight")
		return
	}

	favoritedModel.FavoritedDescription = description
	err = repositories.Save(&favoritedModel)
	return
}

func getAllInsights(insights *[]models.Insight, errs chan error) {
	err := repositories.GetAllInsights(insights)
	errs <- err
}

func getAllInsightsPaginated(insights *[]models.Insight, pageSize uint, offset uint, errs chan error) {
	err := repositories.GetAllInsightsPaginated(insights, pageSize, offset)
	errs <- err
}

func getInsightsByIds(favoritedInsights *[]models.FavoritedInsight, insights *[]models.Insight, errs chan error) {
	insightIds := []uint{}
	for _, insight := range *favoritedInsights {
		insightIds = append(insightIds, insight.InsightID)
	}
	err := repositories.GetInsightsByIds(insightIds, insights)
	errs <- err
}

func getInsightsByIdsPaginated(favoritedInsights *[]models.FavoritedInsight, insights *[]models.Insight, pageSize uint, offset uint, errs chan error) {
	insightIds := []uint{}
	for _, insight := range *favoritedInsights {
		insightIds = append(insightIds, insight.InsightID)
	}
	err := repositories.GetInsightsByIdsPaginated(insightIds, insights, pageSize, offset)
	errs <- err
}