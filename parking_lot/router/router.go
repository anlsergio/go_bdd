package router

import (
	"github.com/anlsergio/go_bdd/parking_lot/router/routes"
	"github.com/gorilla/mux"
)

// Create returns the instance of a new router with all the defined routes registered on it
func Create() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetRouter(r)
}
