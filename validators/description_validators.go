package validators

import (
	"gwi-challenge/common"

	"github.com/gin-gonic/gin"
)

type DescriptionValidator struct {
	DescriptionToValidate string `json:"description" binding:"exists"`
	Description           string `json:"-"`
}

func NewDescriptionValidator() DescriptionValidator {
	return DescriptionValidator{}
}

func (validator *DescriptionValidator) Bind(context *gin.Context) error {
	err := common.Bind(context, validator)
	if err != nil {
		return err
	}

	validator.Description = validator.DescriptionToValidate
	return nil
}
