package main

import (
	"log"
	"net/http"
)

// just a main function that starts the Server
func main() {

	// add a single asset
	fillAssets(1)
	DBaddCreds()

	logWebServer("Server started")

	log.Fatal(http.ListenAndServe(":8080", endPointHandler()))
}
