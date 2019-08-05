package main

import (
	"log"
	"net/http"
)

// just a main function that starts the Server
func main() {
	// create a new DB
	db := CreateNewDB()

	// add a single asset
	db.fillAssets(1)
	// add admin creds
	db.DBaddCreds()

	logWebServer("Server started")

	// start on 8080 while logging each req/resp with LogHTTP
	log.Fatal(http.ListenAndServe(":8080", LogHTTP(endPointHandler(db))))
}
