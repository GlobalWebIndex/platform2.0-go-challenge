package validators

import (
	"gwi-challenge/common"
	"gwi-challenge/data/models"

	"github.com/gin-gonic/gin"
)

type ChartValidator struct {
	ChartToValidate struct {
		Title string                 `json:"title" binding:"exists,min=4"`
		XAxes string                 `json:"x_axes" binding:"exists,min=4"`
		YAxes string                 `json:"y_axes" binding:"exists,min=4"`
		Data  []ChartPointToValidate `json:"data" binding:"exists"`
	} `json:"chart"`
	Chart models.Chart
}

type ChartPointToValidate struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

func NewChartValidator() ChartValidator {
	return ChartValidator{}
}

func (validator *ChartValidator) Bind(context *gin.Context) error {

	// bind json body to our validator
	err := common.Bind(context, validator)
	if err != nil {
		return err
	}

	// bind every property to out validator.chart
	validator.Chart.Title = validator.ChartToValidate.Title
	validator.Chart.XAxes = validator.ChartToValidate.XAxes
	validator.Chart.YAxes = validator.ChartToValidate.YAxes
	validator.Chart.Data = []models.ChartPoint{}

	for _, datum := range validator.ChartToValidate.Data {
		validator.Chart.Data = append(validator.Chart.Data, models.ChartPoint{X: datum.X, Y: datum.Y})
	}

	return nil
}
