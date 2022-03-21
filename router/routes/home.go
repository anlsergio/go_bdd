package routes

import (
	"net/http"

	"github.com/anlsergio/go_bdd/parking_lot/controller"
)

var homeRoutes = []Route{
	{
		URI:      "/",
		Method:   http.MethodGet,
		Function: controller.HomePage,
	},
}
