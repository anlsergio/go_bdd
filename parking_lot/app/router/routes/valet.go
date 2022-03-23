package routes

import (
	"net/http"

	"github.com/anlsergio/go_bdd/parking_lot/controller"
)

var valetRoutes = []Route{
	{
		URI:      "/valet/{minutes}",
		Method:   http.MethodGet,
		Function: controller.CalculateValet,
	},
}
