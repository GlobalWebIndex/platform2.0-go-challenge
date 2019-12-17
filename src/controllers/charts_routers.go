package controllers

import (
	"gwi-challenge/common"
	"gwi-challenge/serializers"
	"gwi-challenge/services"
	"gwi-challenge/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChartsRoutesRegister(router *gin.RouterGroup) {
	router.POST("", createChart)
	router.GET("/:title", getChart)
	router.PUT("/:title", updateChart)
	router.DELETE("/:title", deleteChart)
	router.POST("/:title/favorite", favoriteChart)
	router.DELETE("/:title/favorite", unfavoriteChart)
	router.PATCH("/:title/describe", describeChart)
}

func favoriteChart(context *gin.Context) {
	title := context.Param("title")
	userID := context.GetInt("user_id")
	if err := services.FavoriteChart(uint(userID), title); err != nil {
		context.JSON(http.StatusBadRequest, common.NewError("db_error", err))
		return
	}
}

func unfavoriteChart(context *gin.Context) {
	title := context.Param("title")
	userID := context.GetInt("user_id")
	if err := services.UnfavoriteChart(uint(userID), title); err != nil {
		context.JSON(http.StatusBadRequest, common.NewError("db_error", err))
		return
	}
}

func describeChart(context *gin.Context) {
	title := context.Param("title")
	userID := context.GetInt("user_id")
	validator := validators.NewDescriptionValidator()
	if err := validator.Bind(context); err != nil {
		context.JSON(http.StatusBadRequest, common.NewError("validation_error", err))
		return
	}
	if err := services.DescribeChart(uint(userID), title, validator.Description); err != nil {
		context.JSON(http.StatusBadRequest, common.NewError("db_error", err))
		return
	}
}

func getChart(context *gin.Context) {
	chartTitle := context.Param("title")
	chartToReturn, err := services.GetChartByTitle(chartTitle)
	if err != nil {
		context.JSON(http.StatusNotFound, common.NewError("db_error", err))
		return
	}
	serializer := serializers.NewChartSerializer()
	serializer.Chart = chartToReturn
	context.JSON(http.StatusOK, gin.H{"chart": serializer.Response()})
}

func createChart(context *gin.Context) {
	chartValidator := validators.NewChartValidator()

	if err := chartValidator.Bind(context); err != nil {
		context.JSON(http.StatusBadRequest, common.NewError("validation_error", err))
		return
	}

	if err := services.CreateNewChart(&chartValidator.Chart); err != nil {
		context.JSON(http.StatusUnprocessableEntity, common.NewError("db_error", err))
		return
	}

	serializer := serializers.NewChartSerializer()
	serializer.Chart = &chartValidator.Chart
	context.JSON(http.StatusCreated, gin.H{"chart": serializer.Response()})
}

func updateChart(context *gin.Context) {
	chartTitle := context.Param("title")

	chartValidator := validators.NewChartValidator()
	if err := chartValidator.Bind(context); err != nil {
		context.JSON(http.StatusBadRequest, common.NewError("validation_error", err))
		return
	}

	updatedChart, err := services.UpdateChartByTitle(chartTitle, &chartValidator.Chart)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, common.NewError("db_error", err))
		return
	}

	serializer := serializers.NewChartSerializer()
	serializer.Chart = updatedChart
	context.JSON(http.StatusOK, gin.H{"chart": serializer.Response()})
}

func deleteChart(context *gin.Context) {
	chartTitle := context.Param("title")
	err := services.DeleteChartByTitle(chartTitle)
	if err != nil {
		context.JSON(http.StatusNotFound, common.NewError("db_error", err))
		return
	}
	context.JSON(http.StatusOK, gin.H{})
}
