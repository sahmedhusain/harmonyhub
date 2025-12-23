package main

import (
	"fmt"
	server "harmonyhub/src/server"
	"log"
	"net/http"
)

// Main function to start the server
func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	// Serves static files from the assets directory
	http.HandleFunc("/", server.HomePage)
	// Handles requests to the homepage
	http.HandleFunc("/artist", server.ArtistPage)
	// Handles requests to the artist page
	http.HandleFunc("/search", server.SearchHandler)
	// Handles search requests
	fmt.Println("Server running on http://localhost:8080 \nTo stop the server press Ctrl+C")
	// Prints server running message
	log.Fatal(http.ListenAndServe(":8080", nil))
	// Starts the server on port 8080
}
