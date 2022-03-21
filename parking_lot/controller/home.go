package controller

import (
	"fmt"
	"net/http"
)

// HomePage is the root endpoint controller
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world from BDD with Godog\n")
}
