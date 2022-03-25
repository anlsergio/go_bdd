package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents an API route object
type Route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
}

// SetRouter registers all routes into a given router
func SetRouter(r *mux.Router) *mux.Router {
	routes := homeRoutes
	routes = append(routes, valetRoutes...)
	routes = append(routes, shortTermRoutes...)

	for _, route := range routes {
		r.HandleFunc(
			route.URI,
			route.Function,
		).Methods(route.Method)
	}
	return r
}
