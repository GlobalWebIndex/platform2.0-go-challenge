package main

import (
	"gwiapi/app"
	"gwiapi/config"
)

func main() {

	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}
