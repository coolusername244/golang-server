package main

import (
	"fmt"
	"golang-server/handlers"
	"net/http"
)

func main() {
	// Define routes and handlers
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/hello", handlers.HelloHandler)

	// Start the server
	fmt.Println("Server is listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}


