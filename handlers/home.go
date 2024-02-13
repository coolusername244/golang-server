package handlers

import (
	"fmt"
	"net/http"
)

// Handler function for the home route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}