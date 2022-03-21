package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anlsergio/go_bdd/parking_lot/config"
	"github.com/anlsergio/go_bdd/parking_lot/router"
)

func handleRequests() {
	r := router.Create()

	fmt.Println("The API is running on port", config.ServerPort)
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%d", config.ServerPort), r))
}

func main() {
	handleRequests()
}
