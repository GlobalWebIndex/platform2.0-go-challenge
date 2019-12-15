package validators

import (
	"gwi-challenge/common"
	"gwi-challenge/data/models"

	"github.com/gin-gonic/gin"
)

type InsightValidator struct {
	InsightToValidate struct {
		Title string `json:"title" binding:"exists"`
		Text  string `json:"text" binding:"exists"`
	} `json:"insight"`
	Insight models.Insight `json:"-"`
}

func NewInsightValidator() InsightValidator {
	return InsightValidator{}
}

func (validator *InsightValidator) Bind(context *gin.Context) error {

	// bind json body to our validator
	err := common.Bind(context, validator)
	if err != nil {
		return err
	}

	validator.Insight.Title = validator.InsightToValidate.Title
	validator.Insight.Text = validator.InsightToValidate.Text

	return nil
}
