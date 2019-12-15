package services

import (
	"errors"
	"gwi-challenge/data/models"
)

func GetAllFavoriteAssets(userID uint, assets *models.Assets) (err error) {
	user, err := GetUserById(userID, true)
	if err != nil {
		err = errors.New("User does not exist")
		return
	}
	errs := make(chan error)
	awaitCount := 0
	if len(user.FavoriteCharts) > 0 && assets.Charts != nil {
		go getChartsByIds(&user.FavoriteCharts, assets.Charts, errs)
		awaitCount++
	}
	if len(user.FavoriteInsights) > 0 && assets.Insights != nil {
		go getInsightsByIds(&user.FavoriteInsights, assets.Insights, errs)
		awaitCount++
	}
	if len(user.FavoriteAudiences) > 0 && assets.Audiences != nil {
		go getAudiencesByIds(&user.FavoriteAudiences, assets.Audiences, errs)
		awaitCount++
	}
	for index := 0; index < awaitCount; index++ {
		tempErr := <-errs
		if tempErr != nil {
			err = tempErr
		}
	}
	return
}

func GetAllFavoriteAssetsPaginated(userID uint, assets *models.Assets, pageSize uint, pageNumber uint) (err error) {
	user, err := GetUserById(userID, true)
	if err != nil {
		err = errors.New("User does not exist")
		return
	}
	errs := make(chan error)
	offset := pageNumber*pageSize - pageSize
	awaitCount := 0
	if len(user.FavoriteCharts) > 0 && assets.Charts != nil {
		go getChartsByIdsPaginated(&user.FavoriteCharts, assets.Charts, pageSize, offset, errs)
		awaitCount++
	}
	if len(user.FavoriteInsights) > 0 && assets.Insights != nil {
		go getInsightsByIdsPaginated(&user.FavoriteInsights, assets.Insights, pageSize, offset, errs)
		awaitCount++
	}
	if len(user.FavoriteAudiences) > 0 && assets.Audiences != nil {
		go getAudiencesByIdsPaginated(&user.FavoriteAudiences, assets.Audiences, pageSize, offset, errs)
		awaitCount++
	}
	for index := 0; index < awaitCount; index++ {
		tempErr := <-errs
		if tempErr != nil {
			err = tempErr
		}
	}
	return
}

func GetAllAssets(assets *models.Assets) (err error) {
	errs := make(chan error)

	go getAllCharts(assets.Charts, errs)
	go getAllInsights(assets.Insights, errs)
	go getAllAudiences(assets.Audiences, errs)

	for index := 0; index < 3; index++ {
		tempErr := <-errs
		if tempErr != nil {
			err = tempErr
		}
	}
	return
}

func GetAllAssetsPaginated(assets *models.Assets, pageSize uint, pageNumber uint) (err error) {
	errs := make(chan error)
	offset := pageNumber*pageSize - pageSize

	go getAllChartsPaginated(assets.Charts, pageSize, offset, errs)
	go getAllInsightsPaginated(assets.Insights, pageSize, offset, errs)
	go getAllAudiencesPaginated(assets.Audiences, pageSize, offset, errs)

	for index := 0; index < 3; index++ {
		tempErr := <-errs
		if tempErr != nil {
			err = tempErr
		}
	}
	return
}
