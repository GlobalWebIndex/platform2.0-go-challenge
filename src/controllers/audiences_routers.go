package controllers

import (
	"gwi-challenge/common"
	"gwi-challenge/serializers"
	"gwi-challenge/services"
	"gwi-challenge/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AudiencesRoutesRegister(router *gin.RouterGroup) {
	router.POST("", createAudience)
	router.GET("/:title", getAudience)
	router.DELETE("/:title", deleteAudience)
	router.PUT("/:title", updateAudience)
	router.POST("/:title/favorite", favoriteAudience)
	router.DELETE("/:title/unfavorite", unfavoriteAudience)
	router.PATCH("/:title/describe", describeAudience)
}

func favoriteAudience(context *gin.Context) {
	title := context.Param("title")
	userID := context.GetInt("user_id")
	if err := services.FavoriteAudience(uint(userID), title); err != nil {
		context.JSON(http.StatusBadRequest, common.NewError("db_error", err))
		return
	}
}

func unfavoriteAudience(context *gin.Context) {
	title := context.Param("title")
	userID := context.GetInt("user_id")
	if err := services.UnfavoriteAudience(uint(userID), title); err != nil {
		context.JSON(http.StatusBadRequest, common.NewError("db_error", err))
		return
	}
}

func describeAudience(context *gin.Context) {
	title := context.Param("title")
	userID := context.GetInt("user_id")
	validator := validators.NewDescriptionValidator()
	if err := validator.Bind(context); err != nil {
		context.JSON(http.StatusBadRequest, common.NewError("validation_error", err))
		return
	}
	if err := services.DescribeAudience(uint(userID), title, validator.Description); err != nil {
		context.JSON(http.StatusBadRequest, common.NewError("db_error", err))
		return
	}
}

func getAudience(context *gin.Context) {
	title := context.Param("title")
	audienceToReturn, err := services.GetAudienceByTitle(title)
	if err != nil {
		context.JSON(http.StatusNotFound, common.NewError("db_error", err))
		return
	}

	serializer := serializers.NewAudienceSerializer()
	serializer.Audience = audienceToReturn
	context.JSON(http.StatusOK, gin.H{"audience": serializer.Response()})
}

func deleteAudience(context *gin.Context) {
	title := context.Param("title")
	err := services.DeleteAudienceByTitle(title)
	if err != nil {
		context.JSON(http.StatusNotFound, common.NewError("db_error", err))
		return
	}
	context.JSON(http.StatusOK, "")
}

func createAudience(context *gin.Context) {
	validator := validators.NewAudienceValidator()

	if err := validator.Bind(context); err != nil {
		context.JSON(http.StatusBadRequest, common.NewError("validation_error", err))
		return
	}

	if err := services.CreateNewAudience(&validator.Audience); err != nil {
		context.JSON(http.StatusUnprocessableEntity, common.NewError("db_error", err))
		return
	}

	serializer := serializers.NewAudienceSerializer()
	serializer.Audience = &validator.Audience
	context.JSON(http.StatusCreated, gin.H{"audience": serializer.Response()})
}

func updateAudience(context *gin.Context) {
	title := context.Param("title")

	validator := validators.NewAudienceValidator()

	if err := validator.Bind(context); err != nil {
		context.JSON(http.StatusBadRequest, common.NewError("validation_error", err))
		return
	}

	updateAudience, err := services.UpdateAudienceByTitle(title, &validator.Audience)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, common.NewError("db_error", err))
		return
	}

	serializer := serializers.NewAudienceSerializer()
	serializer.Audience = updateAudience
	context.JSON(http.StatusOK, gin.H{"audience": serializer.Response()})
}
