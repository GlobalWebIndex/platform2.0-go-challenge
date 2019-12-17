package validators

import (
	"errors"
	"gwi-challenge/common"
	"gwi-challenge/data/models"

	"github.com/gin-gonic/gin"
)

type AudienceValidator struct {
	AudieceToValidate struct {
		Title                   string        `json:"title" binding:"exists,min=4"`
		BirthCountry            string        `json:"birth_country" binding:"exists,min=4"`
		Gender                  models.Gender `json:"gender" binding:"exists"`
		AgeGroupLowerLimit      uint          `json:"lower_age_limit" binding:"exists"`
		AgeGroupUpperLimit      uint          `json:"upper_age_limit" binding:"exists"`
		HoursSpentOnSocialMedia uint          `json:"social_media_hours"`
		PurchasesPerMonth       uint          `json:"monthly_purchases"`
	} `json:"audience"`
	Audience models.Audience
}

func NewAudienceValidator() AudienceValidator {
	return AudienceValidator{}
}

func (validator *AudienceValidator) Bind(context *gin.Context) error {

	err := common.Bind(context, validator)
	if err != nil {
		return err
	}

	validator.Audience.Title = validator.AudieceToValidate.Title
	validator.Audience.BirthCountry = validator.AudieceToValidate.BirthCountry
	validator.Audience.AgeGroupLowerLimit = validator.AudieceToValidate.AgeGroupLowerLimit
	validator.Audience.AgeGroupUpperLimit = validator.AudieceToValidate.AgeGroupUpperLimit

	if validator.AudieceToValidate.Gender.Valid() {
		validator.Audience.Gender = validator.AudieceToValidate.Gender
	} else {
		return errors.New("gender is not valid")
	}

	if validator.AudieceToValidate.HoursSpentOnSocialMedia == 0 &&
		validator.AudieceToValidate.PurchasesPerMonth == 0 {
		return errors.New("you have to provide one audience info")
	}

	if validator.AudieceToValidate.HoursSpentOnSocialMedia > 0 && validator.AudieceToValidate.PurchasesPerMonth == 0 {
		validator.Audience.AudienceInfo = &models.AudienceInfo{
			AudienceInfoType: models.HoursOnSocialMedia,
			AudienceInfoStat: validator.AudieceToValidate.HoursSpentOnSocialMedia,
		}
		return nil
	}

	if validator.AudieceToValidate.PurchasesPerMonth > 0 && validator.AudieceToValidate.HoursSpentOnSocialMedia == 0 {
		validator.Audience.AudienceInfo = &models.AudienceInfo{
			AudienceInfoType: models.PurchasesPerMonth,
			AudienceInfoStat: validator.AudieceToValidate.PurchasesPerMonth,
		}
		return nil
	}

	return errors.New("you have to provide only one audience info")
}
