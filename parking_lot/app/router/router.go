package router

import (
	"github.com/anlsergio/go_bdd/parking_lot/app/router/routes"
	"github.com/gorilla/mux"
)

// New returns the instance of a new router with all the defined routes registered on it
func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetRouter(r)
}
