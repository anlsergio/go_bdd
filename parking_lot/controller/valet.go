package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// CalculateValet calculates the parking cost for Valet parking
func CalculateValet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	minutesSpent := vars["minutes"]
	fmt.Println("minutes spent at the parking a lot: ", minutesSpent)
	fmt.Fprintf(w, "$ 12.00")
}
