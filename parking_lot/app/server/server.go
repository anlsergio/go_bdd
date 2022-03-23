package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anlsergio/go_bdd/parking_lot/app/router"
	"github.com/anlsergio/go_bdd/parking_lot/config"
	"github.com/gorilla/mux"
)

// Server represents the core API server object
type Server struct {
	Router *mux.Router
}

// New returns an instance of the App Server
func New() *Server {
	return &Server{
		Router: router.New(),
	}
}

// Run gets the server ready to listen for requests
func (s *Server) Run(addr string) {
	fmt.Println("The API is running on port", config.ServerPort)

	log.Fatal(http.ListenAndServe(addr, s.Router))
}
