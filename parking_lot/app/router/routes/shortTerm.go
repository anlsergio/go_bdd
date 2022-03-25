package routes

import (
	"net/http"

	"github.com/anlsergio/go_bdd/parking_lot/controller"
)

var shortTermRoutes = []Route{
	{
		URI:      "/short-term/{minutes}",
		Method:   http.MethodGet,
		Function: controller.CalculateShortTerm,
	},
}
