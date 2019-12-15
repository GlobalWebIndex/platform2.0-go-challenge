package validators

import (
	"errors"
	"gwi-challenge/common"
	"gwi-challenge/data/models"
	"regexp"

	"github.com/gin-gonic/gin"
)

type UserRegisterValidator struct {
	UserToValidate struct {
		Username         string `json:"username" binding:"exists,min=6"`
		Password         string `json:"password" binding:"exists,min=6"`
		RepeatedPassword string `json:"password_repeat" binding:"exists,min=6"`
		Name             string `json:"name" binding:"exists,min=6"`
	} `json:"user"`
	User models.User
}

type UserLoginValidator struct {
	UserLoginRequestToValidate struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"user"`
}

var containsAtLeastOneUpperCaseLetter = regexp.MustCompile(`[A-Z]`)
var containsAtLeastOneLowerCaseLetter = regexp.MustCompile(`[a-z]`)
var containsAtLeastOneNumber = regexp.MustCompile(`[0-9]`)
var containsAtLeastOneSpecialCharacter = regexp.MustCompile(`[%#!$^&*@]`)

// NewUserRegisterValidator returns an empty register validator
func NewUserRegisterValidator() UserRegisterValidator {
	return UserRegisterValidator{}
}

// NewUserLoginValidator returns an empty login validator
func NewUserLoginValidator() UserLoginValidator {
	return UserLoginValidator{}
}

// Bind converts the json request to our struct
func (validator *UserRegisterValidator) Bind(context *gin.Context) error {
	err := common.Bind(context, validator)
	if err != nil {
		return err
	}

	if validator.UserToValidate.RepeatedPassword != validator.UserToValidate.Password {
		return errors.New("given passwords must be the same")
	}

	if !containsAtLeastOneUpperCaseLetter.MatchString(validator.UserToValidate.Password) {
		return errors.New("password must contain an upper-case letter")
	}

	if !containsAtLeastOneLowerCaseLetter.MatchString(validator.UserToValidate.Password) {
		return errors.New("password must contain a lower-case letter")
	}

	if !containsAtLeastOneNumber.MatchString(validator.UserToValidate.Password) {
		return errors.New("password must contain at least one number")
	}

	if !containsAtLeastOneSpecialCharacter.MatchString(validator.UserToValidate.Password) {
		return errors.New("password must contain at least one special character")
	}

	validator.User.Username = validator.UserToValidate.Username
	validator.User.FullName = validator.UserToValidate.Name
	validator.User.PasswordHash = validator.UserToValidate.Password

	return nil
}

// Bind converts the json request to our struct
func (validator *UserLoginValidator) Bind(context *gin.Context) error {
	err := common.Bind(context, validator)
	if err != nil {
		return err
	}
	return nil
}
