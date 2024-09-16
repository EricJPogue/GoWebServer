package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type to HTML
	w.Header().Set("Content-Type", "text/html")
	// Write the HTML response
	fmt.Fprintf(w, "<h5>Hello, World! - Version 2</h5>")
}

func main() {
	// Route for the root path
	http.HandleFunc("/", helloWorld)

	// Start the server on port 8080
	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
