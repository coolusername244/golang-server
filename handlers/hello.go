package handlers

import (
	"fmt"
	"net/http"
)

// Handler function for the /hello route
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
