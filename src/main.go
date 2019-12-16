package main

import (
	"fmt"
	"gwi-challenge/common"
	"gwi-challenge/controllers"
	"gwi-challenge/data"
	"gwi-challenge/data/models"
	"gwi-challenge/services"

	"github.com/gin-gonic/gin"
)

func main() {

	common.InitConfig()
	data.InitDB()

	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	gwiAPI := router.Group("/api")

	usersGroup := gwiAPI.Group("/users")
	controllers.UsersRoutesRegister(usersGroup)

	assetsGroup := gwiAPI.Group("/assets")
	assetsGroup.Use(controllers.AuthMiddleware(true))
	controllers.AssetsRoutesRegister(assetsGroup)

	chartsGroup := assetsGroup.Group("/chart")
	controllers.ChartsRoutesRegister(chartsGroup)

	insightsGroup := assetsGroup.Group("/insight")
	controllers.InsightsRoutesRegister(insightsGroup)

	audiencesGroup := assetsGroup.Group("/audience")
	controllers.AudiencesRoutesRegister(audiencesGroup)

	if common.IsDemo() {
		go setDataForDemoPurposes()
	}

	router.Run()
}

func setDataForDemoPurposes() {
	// DB population with demo data
	data.PopulateDbWithDemoData()

	// Demo user registration
	demoUserToRegister := models.User{
		Username:     "demouser",
		FullName:     "Demo User",
		PasswordHash: "password123",
	}
	err := services.CreateNewUser(&demoUserToRegister)

	if err != nil {
		fmt.Println("Something went wrnog with demo user registration :(")
		return
	}

	fmt.Println("Population of db with demo data completed successfully! :)")
}
