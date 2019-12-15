package controllers

import (
	"gwi-challenge/common"
	"gwi-challenge/serializers"
	"gwi-challenge/services"
	"gwi-challenge/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AssetsRoutesRegister(router *gin.RouterGroup) {
	router.GET("/", getAllPaginated)
	router.GET("/all", getAll)
	router.GET("/favorites", getAllFavoritesPaginated)
	router.GET("/favorites/all", getAllFavorites)
}

func getAll(context *gin.Context) {
	serializer := serializers.NewAssetSerializer()
	if err := services.GetAllAssets(&serializer.Assets); err != nil {
		context.JSON(http.StatusUnprocessableEntity, common.NewError("db_error", err))
		return
	}
	context.JSON(http.StatusOK, gin.H{"assets": serializer.Response()})
}

func getAllPaginated(context *gin.Context) {
	validator := validators.NewPagingValidator()
	pageSize, pageNumber, err := validator.GetPagingProps(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, common.NewError("validation_error", err))
		return
	}
	serializer := serializers.NewAssetSerializer()
	if err := services.GetAllAssetsPaginated(&serializer.Assets, uint(pageSize), uint(pageNumber)); err != nil {
		context.JSON(http.StatusUnprocessableEntity, common.NewError("db_error", err))
		return
	}
	context.JSON(http.StatusOK, gin.H{"assets": serializer.Response()})
}

func getAllFavorites(context *gin.Context) {
	userID := context.GetInt("user_id")
	serializer := serializers.NewAssetSerializer()
	if err := services.GetAllFavoriteAssets(uint(userID), &serializer.Assets); err != nil {
		context.JSON(http.StatusUnprocessableEntity, common.NewError("db_error", err))
		return
	}
	context.JSON(http.StatusOK, gin.H{"favorited_assets": serializer.Response()})
}

func getAllFavoritesPaginated(context *gin.Context) {
	userID := context.GetInt("user_id")
	validator := validators.NewPagingValidator()
	pageSize, pageNumber, err := validator.GetPagingProps(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, common.NewError("validation_error", err))
		return
	}
	serializer := serializers.NewAssetSerializer()
	if err := services.GetAllFavoriteAssetsPaginated(uint(userID), &serializer.Assets, uint(pageSize), uint(pageNumber)); err != nil {
		context.JSON(http.StatusUnprocessableEntity, common.NewError("db_error", err))
		return
	}
	context.JSON(http.StatusOK, gin.H{"favorited_assets": serializer.Response()})
}
