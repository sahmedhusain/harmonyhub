package server // Define the package name

import (
	"encoding/json" // Import the JSON package for encoding and decoding
	"fmt"           // Import the fmt package for formatted I/O
	"html/template" // Import the template package for HTML templates
	"log"           // Import the log package for logging
	"net/http"      // Import the net/http package for HTTP client and server
	"strings"       // Import the strings package for string manipulation
	"sync"          // Import the sync package for synchronization
)

// Artist struct to hold the artist data
type Artist struct {
	ID           int      `json:"id"`           // ID of the artist
	Name         string   `json:"name"`         // Name of the artist
	Image        string   `json:"image"`        // Image URL of the artist
	Members      []string `json:"members"`      // List of members in the band
	CreationDate int      `json:"creationDate"` // Creation date of the band
	FirstAlbum   string   `json:"firstAlbum"`   // First album of the band
	LocationsURL string   `json:"locations"`    // URL for concert locations
	DatesURL     string   `json:"concertDates"` // URL for concert dates
	RelationsURL string   `json:"relations"`    // URL for relations data
}

// Dates struct to hold the concert dates
type Dates struct {
	ID    int      `json:"id"`    // ID of the date entry
	Dates []string `json:"dates"` // List of concert dates
}

// Locations struct to hold the concert locations
type Locations struct {
	ID        int      `json:"id"`        // ID of the location entry
	Locations []string `json:"locations"` // List of concert locations
}

// Relation struct to hold the relation data
type Relation struct {
	ID             int                 `json:"id"`             // ID of the relation entry
	DatesLocations map[string][]string `json:"datesLocations"` // Map of locations to dates
}

// Data struct to hold all the fetched data
type Data struct {
	Artists   []Artist         // List of artists
	Dates     map[int][]string // Map of artist IDs to concert dates
	Locations map[int][]string // Map of artist IDs to concert locations
	Relations map[int]Relation // Map of artist IDs to relation data
}

type ErrorPageData struct {
	Code     string // Error code (e.g., "404", "500")
	ErrorMsg string // Error message (e.g., "PAGE NOT FOUND", "INTERNAL SERVER ERROR")
}

var data Data      // Global variable to hold the fetched data
var once sync.Once // Synchronization primitive to ensure fetchData is called only once

// errHandler handles errors by rendering the error page
func errHandler(w http.ResponseWriter, r *http.Request, err *ErrorPageData) {
	errorTemp, erra := template.ParseFiles("templates/error.html") // Parse the error HTML template
	if erra != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError) // Respond with internal server error
		return
	}
	errorTemp.Execute(w, err) // Execute the template with the error data
}

// fetchData fetches data from the API
func fetchData() {
	// Fetch artists data
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists") // Send HTTP GET request to fetch artists
	if err != nil {
		log.Fatalf("Failed to fetch artists data: %v", err) // Log fatal error if the request fails
	}
	defer resp.Body.Close() // Ensure the response body is closed

	err = json.NewDecoder(resp.Body).Decode(&data.Artists) // Decode the response body into the artists slice
	if err != nil {
		log.Fatalf("Failed to decode artists data: %v", err) // Log fatal error if decoding fails
	}

	// Fetch dates data
	resp, err = http.Get("https://groupietrackers.herokuapp.com/api/dates") // Send HTTP GET request to fetch dates
	if err != nil {
		log.Fatalf("Failed to fetch dates data: %v", err) // Log fatal error if the request fails
	}
	defer resp.Body.Close() // Ensure the response body is closed

	var datesResponse struct {
		Index []Dates `json:"index"` // Define a struct to hold the dates response
	}
	err = json.NewDecoder(resp.Body).Decode(&datesResponse) // Decode the response body into the datesResponse struct
	if err != nil {
		log.Fatalf("Failed to decode dates data: %v", err) // Log fatal error if decoding fails
	}
	data.Dates = make(map[int][]string) // Initialize the Dates map
	for _, date := range datesResponse.Index {
		data.Dates[date.ID] = date.Dates // Populate the Dates map with IDs and dates
	}

	// Fetch locations data
	resp, err = http.Get("https://groupietrackers.herokuapp.com/api/locations") // Send HTTP GET request to fetch locations
	if err != nil {
		log.Fatalf("Failed to fetch locations data: %v", err) // Log fatal error if the request fails
	}
	defer resp.Body.Close() // Ensure the response body is closed

	var locationsResponse struct {
		Index []Locations `json:"index"` // Define a struct to hold the locations response
	}
	err = json.NewDecoder(resp.Body).Decode(&locationsResponse) // Decode the response body into the locationsResponse struct
	if err != nil {
		log.Fatalf("Failed to decode locations data: %v", err) // Log fatal error if decoding fails
	}
	data.Locations = make(map[int][]string) // Initialize the Locations map
	for _, location := range locationsResponse.Index {
		data.Locations[location.ID] = location.Locations // Populate the Locations map with IDs and locations
	}

	// Fetch relations data
	resp, err = http.Get("https://groupietrackers.herokuapp.com/api/relation") // Send HTTP GET request to fetch relations
	if err != nil {
		log.Fatalf("Failed to fetch relations data: %v", err) // Log fatal error if the request fails
	}
	defer resp.Body.Close() // Ensure the response body is closed

	var relationsResponse struct {
		Index []Relation `json:"index"` // Define a struct to hold the relations response
	}
	err = json.NewDecoder(resp.Body).Decode(&relationsResponse) // Decode the response body into the relationsResponse struct
	if err != nil {
		log.Fatalf("Failed to decode relations data: %v", err) // Log fatal error if decoding fails
	}
	data.Relations = make(map[int]Relation) // Initialize the Relations map
	for _, relation := range relationsResponse.Index {
		data.Relations[relation.ID] = relation // Populate the Relations map with IDs and relations
	}
}

// clean removes unwanted characters from a string
func clean(s string) string {
	s = strings.ReplaceAll(s, "*", "")  // Replace asterisks with empty strings
	s = strings.ReplaceAll(s, "_", " ") // Replace underscores with spaces
	return s
}

// HomePage handler for the home page
func HomePage(w http.ResponseWriter, r *http.Request) {
	// Validating the request path
	if r.URL.Path != "/" { // Check if the request path is not the root
		err := ErrorPageData{Code: "404", ErrorMsg: "PAGE NOT FOUND"} // Create error data for "PAGE NOT FOUND"
		w.WriteHeader(http.StatusNotFound)                            // Set HTTP status to 404
		errHandler(w, r, &err)                                        // Render the error page
		return
	}
	// Validating the request method
	if r.Method != "GET" { // Check if the request method is not GET
		err := ErrorPageData{Code: "405", ErrorMsg: "METHOD NOT ALLOWED"} // Create error data for "METHOD NOT ALLOWED"
		w.WriteHeader(http.StatusMethodNotAllowed)                        // Set HTTP status to 405
		errHandler(w, r, &err)                                            // Render the error page
		return
	}
	// Validating the parsing of the main page
	main, err := template.ParseFiles("templates/index.html") // Parse the main HTML template
	if err != nil {                                          // Check if there was an error parsing the template
		err := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"} // Create error data for "INTERNAL SERVER ERROR"
		w.WriteHeader(http.StatusInternalServerError)                        // Set HTTP status to 500
		errHandler(w, r, &err)                                               // Render the error page
		return
	}

	once.Do(fetchData) // Fetch data only once

	// Execute the template with the fetched data
	err = main.Execute(w, data.Artists)
	if err != nil {
		log.Fatalf("Failed to execute template: %v", err) // Log fatal error if template execution fails
	}
}

// ArtistPage handler for artist details page
func ArtistPage(w http.ResponseWriter, r *http.Request) {
	// Validating the request path
	id := r.URL.Query().Get("id") // Get the artist ID from the query parameters
	if id == "" {                 // Check if the artist ID is missing
		err := ErrorPageData{Code: "400", ErrorMsg: "BAD REQUEST"} // Create error data for "BAD REQUEST"
		w.WriteHeader(http.StatusBadRequest)                       // Set HTTP status to 400
		errHandler(w, r, &err)                                     // Render the error page
		return
	}

	// Validating the parsing of the artist page
	pg, err := template.ParseFiles("templates/artist.html") // Parse the artist HTML template
	if err != nil {                                         // Check if there was an error parsing the template
		err := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"} // Create error data for “INTERNAL SERVER ERROR”
		w.WriteHeader(http.StatusInternalServerError)                        // Set HTTP status to 500
		errHandler(w, r, &err)                                               // Render the error page
		return
	}
	once.Do(fetchData)

	// Find the artist by ID
	var artist Artist
	for _, a := range data.Artists {
		if fmt.Sprintf("%d", a.ID) == id {
			artist = a
			break
		}
	}
	if artist.ID == 0 { // Check if the artist was not found
		err := ErrorPageData{Code: "404", ErrorMsg: "ARTIST NOT FOUND"} // Create error data for "ARTIST NOT FOUND"
		w.WriteHeader(http.StatusNotFound)                              // Set HTTP status to 404
		errHandler(w, r, &err)                                          // Render the error page
		return
	}

	// Add concert dates and locations to the artist
	artistConcertDates := data.Dates[artist.ID]
	artistConcertLocations := data.Locations[artist.ID]
	//apply the clean func on locations & dates
	for i, loc := range artistConcertLocations {
		artistConcertLocations[i] = clean(loc)
	}

	for i, date := range artistConcertDates {
		artistConcertDates[i] = clean(date)
	}

	// link concert dates with concert location
	linkedConcerts := make(map[string][]string)
	for location, dates := range data.Relations[artist.ID].DatesLocations {
		for _, date := range dates {
			linkedConcerts[clean(location)] = append(linkedConcerts[clean(location)], clean(date))
		}
	}

	// Format members string
	cleanedMembers := make([]string, len(artist.Members))
	for i, member := range artist.Members {
		cleanedMembers[i] = clean(strings.TrimSpace(member))
	}
	members := strings.Join(cleanedMembers, ", ")
	err = pg.Execute(w, map[string]interface{}{
		"Artist":         artist,
		"Members":        members,
		"LinkedConcerts": linkedConcerts,
	})
	if err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}
}

// SearchHandler handles the search requests
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	query = strings.ToLower(query)
	suggestions := []SearchItem{}

	// Ensure data is fetched
	once.Do(fetchData)

	for _, artist := range data.Artists {
		if strings.Contains(strings.ToLower(artist.Name), query) {
			suggestions = append(suggestions, SearchItem{Name: artist.Name, Type: "artist/band", ID: artist.ID})
		}
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), query) {
				suggestions = append(suggestions, SearchItem{Name: member, Type: "member", ID: artist.ID})
			}
		}
		if strings.Contains(fmt.Sprintf("%d", artist.CreationDate), query) {
			suggestions = append(suggestions, SearchItem{Name: fmt.Sprintf("%d", artist.CreationDate), Type: "creation date", ID: artist.ID})
		}
		if strings.Contains(strings.ToLower(artist.FirstAlbum), query) {
			suggestions = append(suggestions, SearchItem{Name: artist.FirstAlbum, Type: "first album", ID: artist.ID})
		}

		// Search in concert locations
		for _, location := range data.Locations[artist.ID] {
			if strings.Contains(strings.ToLower(location), query) {
				suggestions = append(suggestions, SearchItem{Name: location, Type: "location", ID: artist.ID})
			}
		}
	}

	response, err := json.Marshal(suggestions)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// SearchItem struct to hold search suggestions
type SearchItem struct {
	Name string `json:"name"`
	Type string `json:"type"`
	ID   int    `json:"id"`
}
