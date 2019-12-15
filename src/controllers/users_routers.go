package controllers

import (
	"gwi-challenge/common"
	"gwi-challenge/serializers"
	"gwi-challenge/services"
	"gwi-challenge/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UsersRoutesRegister registers all user related endpoints
func UsersRoutesRegister(router *gin.RouterGroup) {
	router.POST("/register", registerUser)
	router.POST("/login", loginUser)
}

func registerUser(context *gin.Context) {
	validator := validators.NewUserRegisterValidator()

	if err := validator.Bind(context); err != nil {
		context.JSON(http.StatusBadRequest, common.NewError("validation_error", err))
		return
	}

	if err := services.CreateNewUser(&validator.User); err != nil {
		context.JSON(http.StatusUnprocessableEntity, common.NewError("db_error", err))
		return
	}

	context.JSON(http.StatusCreated, gin.H{})
}

func loginUser(context *gin.Context) {
	validator := validators.NewUserLoginValidator()

	if err := validator.Bind(context); err != nil {
		context.JSON(http.StatusBadRequest, common.NewError("validation_error", err))
		return
	}

	user, accessToken, err := services.GetUserWithAccessToken(validator.UserLoginRequestToValidate.Username, validator.UserLoginRequestToValidate.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, common.NewError("login_error", err))
	}

	serializer := serializers.NewLoginUserSerializer()
	serializer.User = user
	serializer.AccessToken = accessToken
	context.JSON(http.StatusOK, gin.H{"chart": serializer.Response()})
}
