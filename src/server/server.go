package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"sync"
)

// Artist struct to hold the artist data
type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	LocationsURL string   `json:"locations"`
	DatesURL     string   `json:"concertDates"`
	RelationsURL string   `json:"relations"`
}

// Dates struct to hold the concert dates
type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// Locations struct to hold the concert locations
type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

// Relation struct to hold the relation data
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Data struct to hold all the fetched data
type Data struct {
	Artists   []Artist
	Dates     map[int][]string
	Locations map[int][]string
	Relations map[int]Relation
}

type ErrorPageData struct {
	Code     string // Error code (e.g., "404", "500")
	ErrorMsg string // Error message (e.g., "PAGE NOT FOUND", "INTERNAL SERVER ERROR")
}

var data Data
var once sync.Once

func errHandler(w http.ResponseWriter, r *http.Request, err *ErrorPageData) {
	errorTemp, erra := template.ParseFiles("templates/error.html") // Parse the error HTML template
	if erra != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	errorTemp.Execute(w, err) // Execute the template with the error data
}

// Fetch data from the API
func fetchData() {
	// Fetch artists data
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatalf("Failed to fetch artists data: %v", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&data.Artists)
	if err != nil {
		log.Fatalf("Failed to decode artists data: %v", err)
	}

	// Fetch dates data
	resp, err = http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		log.Fatalf("Failed to fetch dates data: %v", err)
	}
	defer resp.Body.Close()

	var datesResponse struct {
		Index []Dates `json:"index"`
	}
	err = json.NewDecoder(resp.Body).Decode(&datesResponse)
	if err != nil {
		log.Fatalf("Failed to decode dates data: %v", err)
	}
	data.Dates = make(map[int][]string)
	for _, date := range datesResponse.Index {
		data.Dates[date.ID] = date.Dates
	}

	// Fetch locations data
	resp, err = http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		log.Fatalf("Failed to fetch locations data: %v", err)
	}
	defer resp.Body.Close()

	var locationsResponse struct {
		Index []Locations `json:"index"`
	}
	err = json.NewDecoder(resp.Body).Decode(&locationsResponse)
	if err != nil {
		log.Fatalf("Failed to decode locations data: %v", err)
	}
	data.Locations = make(map[int][]string)
	for _, location := range locationsResponse.Index {
		data.Locations[location.ID] = location.Locations
	}

	// Fetch relations data
	resp, err = http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		log.Fatalf("Failed to fetch relations data: %v", err)
	}
	defer resp.Body.Close()

	var relationsResponse struct {
		Index []Relation `json:"index"`
	}
	err = json.NewDecoder(resp.Body).Decode(&relationsResponse)
	if err != nil {
		log.Fatalf("Failed to decode relations data: %v", err)
	}
	data.Relations = make(map[int]Relation)
	for _, relation := range relationsResponse.Index {
		data.Relations[relation.ID] = relation
	}
}

// remove unwanted strings
func clean(s string) string {
	s = strings.ReplaceAll(s, "*", "")
	s = strings.ReplaceAll(s, "_", " ")
	return s
}

// Handler for the home page
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

	once.Do(fetchData)

	// Execute the template with the fetched data
	err = main.Execute(w, data.Artists)
	if err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}
}

// Handler for artist details page
func ArtistPage(w http.ResponseWriter, r *http.Request) {
	// Validating the request path
	id := r.URL.Query().Get("id")
	if id == "" { // Check if the artist ID is missing
		err := ErrorPageData{Code: "400", ErrorMsg: "BAD REQUEST"} // Create error data for "BAD REQUEST"
		w.WriteHeader(http.StatusBadRequest)                       // Set HTTP status to 400
		errHandler(w, r, &err)                                     // Render the error page
		return
	}

	// Validating the parsing of the artist page
	pg, err := template.ParseFiles("templates/artist.html") // Parse the HTML template
	if err != nil {                                         // Check if there was an error parsing the template
		err := ErrorPageData{Code: "500", ErrorMsg: "INTERNAL SERVER ERROR"} // Create error data for "INTERNAL SERVER ERROR"
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
