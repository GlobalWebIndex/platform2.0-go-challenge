package main

import "log"

func main() {
	log.Printf("Attempting to create server at port %d", port)
	err := handleRequests()
	if err != nil {
		log.Fatal(err)
	}
}
