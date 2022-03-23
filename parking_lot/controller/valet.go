package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/anlsergio/go_bdd/parking_lot/domain/valet"
	"github.com/gorilla/mux"
)

// CalculateValet calculates the parking cost for Valet parking
func CalculateValet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	minutesSpent, _ := strconv.Atoi(vars["minutes"])

	fmt.Fprintf(
		w,
		"$%.2f",
		valet.New().CalculateParkingCost(minutesSpent),
	)
}
