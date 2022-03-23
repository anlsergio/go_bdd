package main

import (
	"fmt"

	"github.com/anlsergio/go_bdd/parking_lot/app/server"
	"github.com/anlsergio/go_bdd/parking_lot/config"
)

func handleRequests() {
	s := server.New()
	s.Run(fmt.Sprintf(":%d", config.ServerPort))
}

func main() {
	handleRequests()
}
