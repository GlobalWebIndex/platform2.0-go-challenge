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

	// start on 8080 while logging each req/resp with LogHTTP
	log.Fatal(http.ListenAndServe(":8080", LogHTTP(endPointHandler())))
}
