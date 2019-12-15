package controllers

import (
	"gwi-challenge/common"
	"gwi-challenge/serializers"
	"gwi-challenge/services"
	"gwi-challenge/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsightsRoutesRegister(router *gin.RouterGroup) {
	router.POST("", createInsight)
	router.GET("/:title", getInsight)
	router.PUT("/:title", updateInsight)
	router.DELETE("/:title", deleteInsight)
	router.POST("/:title/favorite", favoriteInsight)
	router.DELETE("/:title/unfavorite", unfavoriteInsight)
	router.PATCH("/:title/describe", describeInsight)
}

func favoriteInsight(context *gin.Context) {
	title := context.Param("title")
	userID := context.GetInt("user_id")
	if err := services.FavoriteInsight(uint(userID), title); err != nil {
		context.JSON(http.StatusBadRequest, common.NewError("db_error", err))
		return
	}
}

func unfavoriteInsight(context *gin.Context) {
	title := context.Param("title")
	userID := context.GetInt("user_id")
	if err := services.UnfavoriteInsight(uint(userID), title); err != nil {
		context.JSON(http.StatusBadRequest, common.NewError("db_error", err))
		return
	}
}

func describeInsight(context *gin.Context) {
	title := context.Param("title")
	userID := context.GetInt("user_id")
	validator := validators.NewDescriptionValidator()
	if err := validator.Bind(context); err != nil {
		context.JSON(http.StatusBadRequest, common.NewError("validation_error", err))
		return
	}
	if err := services.DescribeInsight(uint(userID), title, validator.Description); err != nil {
		context.JSON(http.StatusBadRequest, common.NewError("db_error", err))
		return
	}
}

func createInsight(context *gin.Context) {
	validator := validators.NewInsightValidator()
	if err := validator.Bind(context); err != nil {
		context.JSON(http.StatusBadRequest, common.NewError("validation_error", err))
		return
	}
	if err := services.CreateNewInsight(&validator.Insight); err != nil {
		context.JSON(http.StatusUnprocessableEntity, common.NewError("db_error", err))
		return
	}
	serializer := serializers.NewInsightSerializer()
	serializer.Insight = &validator.Insight
	context.JSON(http.StatusCreated, gin.H{"insight": serializer.Response()})
}

func getInsight(context *gin.Context) {
	title := context.Param("title")
	insightToReturn, err := services.GetInsightByTitle(title)
	if err != nil {
		context.JSON(http.StatusNotFound, common.NewError("db_error", err))
		return
	}
	serializer := serializers.NewInsightSerializer()
	serializer.Insight = insightToReturn
	context.JSON(http.StatusOK, gin.H{"insight": serializer.Response()})
}

func updateInsight(context *gin.Context) {
	title := context.Param("title")
	validator := validators.NewInsightValidator()
	if err := validator.Bind(context); err != nil {
		context.JSON(http.StatusBadRequest, common.NewError("validation_error", err))
		return
	}
	updatedInsight, err := services.UpdateInsightByTitle(title, &validator.Insight)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, common.NewError("db_error", err))
		return
	}
	serializer := serializers.NewInsightSerializer()
	serializer.Insight = updatedInsight
	context.JSON(http.StatusOK, gin.H{"insight": serializer.Response()})
}

func deleteInsight(context *gin.Context) {
	title := context.Param("title")
	err := services.DeleteInsightByTitle(title)
	if err != nil {
		context.JSON(http.StatusNotFound, common.NewError("db_error", err))
		return
	}
	context.JSON(http.StatusOK, gin.H{})
}
