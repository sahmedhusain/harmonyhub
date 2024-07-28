package main

import (
	"fmt"
	server "groupie-tracker/src/server"
	"log"
	"net/http"
)

// Main function to start the server
func main() {

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", server.HomePage)
	http.HandleFunc("/artist", server.ArtistPage)

	fmt.Println("Server running on http://localhost:8080 \nTo stop the server press Ctrl+C")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
