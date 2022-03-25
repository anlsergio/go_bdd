package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/anlsergio/go_bdd/parking_lot/model/shortterm"
	"github.com/gorilla/mux"
)

// CalculateShortTerm calculates the parking cost for Short Term parking
func CalculateShortTerm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	minutesSpent, _ := strconv.Atoi(vars["minutes"])

	fmt.Fprintf(w, "$%.2f", shortterm.New().CalculateParkingCost(minutesSpent))
}
