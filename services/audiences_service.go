package services

import (
	"errors"
	"gwi-challenge/data/models"
	"gwi-challenge/data/repositories"
	"strings"
)

func GetAudienceByTitle(title string) (audienceToReturn *models.Audience, err error) {
	audienceToReturn, err = repositories.GetAudienceByTitle(title)
	return
}

func CreateNewAudience(audienceToCreate *models.Audience) error {
	err := repositories.Save(&audienceToCreate)
	if err != nil && strings.Contains(err.Error(), "constraint") {
		err = errors.New("title already exists")
	}
	return err
}

func DeleteAudienceByTitle(title string) (err error) {
	audienceToDelete, err := repositories.GetAudienceByTitle(title)
	if err != nil {
		return
	}
	err = repositories.Delete(&audienceToDelete.AudienceInfo)
	if err != nil {
		return
	}
	err = repositories.Delete(&audienceToDelete)
	return
}

func UpdateAudienceByTitle(title string, newAudience *models.Audience) (updatedAudience *models.Audience, err error) {
	updatedAudience, err = repositories.GetAudienceByTitle(title)
	if err != nil {
		return nil, err
	}

	updatedAudience.Title = newAudience.Title
	updatedAudience.BirthCountry = newAudience.BirthCountry
	updatedAudience.AgeGroupLowerLimit = newAudience.AgeGroupLowerLimit
	updatedAudience.AgeGroupUpperLimit = newAudience.AgeGroupUpperLimit
	updatedAudience.Gender = newAudience.Gender
	updatedAudience.AudienceInfo.AudienceInfoType = newAudience.AudienceInfo.AudienceInfoType
	updatedAudience.AudienceInfo.AudienceInfoStat = newAudience.AudienceInfo.AudienceInfoStat
	err = repositories.Save(&updatedAudience)
	if err != nil && strings.Contains(err.Error(), "constraint") {
		err = errors.New("title already exists")
	}
	return
}

func FavoriteAudience(userID uint, title string) (err error) {
	audienceToFavorite, err := GetAudienceByTitle(title)
	if err != nil {
		err = errors.New("Audience with title " + title + " does not exist")
		return
	}
	userWhoLoves, err := GetUserById(userID, false)
	if err != nil {
		err = errors.New("User does not exist")
		return
	}
	favoriteModel := models.FavoritedAudience{
		AudienceID: audienceToFavorite.ID,
		UserID:     userWhoLoves.ID,
	}

	err = repositories.Save(&favoriteModel)
	return
}

func UnfavoriteAudience(userID uint, title string) (err error) {
	audienceToUnfavorite, err := GetAudienceByTitle(title)
	if err != nil {
		err = errors.New("Audience with title " + title + " does not exist")
		return
	}
	userWhoHates, err := GetUserById(userID, false)
	if err != nil {
		err = errors.New("User does not exist")
		return
	}

	favoritedModel, err := repositories.GetFavoritedAudience(userWhoHates.ID, audienceToUnfavorite.ID)
	if err != nil {
		err = errors.New("User has not favorited this audience")
		return
	}
	err = repositories.Delete(&favoritedModel)
	return
}

func DescribeAudience(userID uint, title string, description string) (err error) {
	audienceToPatch, err := GetAudienceByTitle(title)
	if err != nil {
		err = errors.New("Audience with title " + title + " does not exist")
		return
	}
	user, err := GetUserById(userID, false)
	if err != nil {
		err = errors.New("User does not exist")
		return
	}

	favoritedModel, err := repositories.GetFavoritedAudience(user.ID, audienceToPatch.ID)
	if err != nil {
		err = errors.New("User has not favorited this audience")
		return
	}

	favoritedModel.FavoritedDescription = description
	err = repositories.Save(&favoritedModel)
	return
}

func getAllAudiences(audiences *[]models.Audience, errs chan error) {
	err := repositories.GetAllAudiences(audiences)
	errs <- err
}

func getAllAudiencesPaginated(audiences *[]models.Audience, pageSize uint, offset uint, errs chan error) {
	err := repositories.GetAllAudiencesPaginated(audiences, pageSize, offset)
	errs <- err
}

func getAudiencesByIds(favoritedAudiences *[]models.FavoritedAudience, audiences *[]models.Audience, errs chan error) {
	audienceIds := []uint{}
	for _, audience := range *favoritedAudiences {
		audienceIds = append(audienceIds, audience.AudienceID)
	}
	err := repositories.GetAudiencesByIds(audienceIds, audiences)
	errs <- err
}

func getAudiencesByIdsPaginated(favoritedAudiences *[]models.FavoritedAudience, audiences *[]models.Audience, pageSize uint, offset uint, errs chan error) {
	audienceIds := []uint{}
	for _, audience := range *favoritedAudiences {
		audienceIds = append(audienceIds, audience.AudienceID)
	}
	err := repositories.GetAudiencesByIdsPaginated(audienceIds, audiences, pageSize, offset)
	errs <- err
}
