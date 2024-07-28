# Groupie-Tracker

Groupie-Tracker is a web application that I developed to provide detailed information about various artists and bands, including their members, concert dates, and locations. The application fetches data from a given API and presents it in a user-friendly format.

## Objectives

The main objective of Groupie-Tracker is to manipulate data from a provided API to create an informative and interactive website. The API consists of four parts:

1. **Artists**: Information about bands and artists, including their names, images, the year they began their activity, the date of their first album, and the members.
2. **Locations**: Details of their last and/or upcoming concert locations.
3. **Dates**: Information about their last and/or upcoming concert dates.
4. **Relations**: Links between artists, dates, and locations.

Using this data, the website displays the bands' information through various data visualizations such as blocks, cards, tables, lists, pages, and graphics. This project also focuses on the creation of events/actions and their visualization, which involves client-server communication for requesting and receiving information.

## Instructions

- The backend is written in Go.
- The site and server are designed to handle errors gracefully and must not crash at any time.
- All pages function correctly, and error handling is a priority.
- The code adheres to good practices, and test files for unit testing are recommended.

## Allowed Packages

Only the standard Go packages are allowed.

## Usage

To run the Groupie-Tracker application:

1. Clone this repository to your local machine.
2. Ensure you have Go installed.
3. Navigate to the project directory in your terminal.
4. Run the following command to start the application:

```sh
go run main.go
```
5.	Open your web browser and go to http://localhost:8080 to access the application.

## Implementation Details

### Data Fetching

The application fetches data from the provided API endpoints and stores it in structured formats. This includes fetching data about artists, concert dates, locations, and relations. The data is fetched and stored in the Data struct, which is then used to populate the web pages.

### Handlers

The application includes the following handlers:

	•	Home Page Handler: Renders the home page with a list of artists.
	•	Artist Page Handler: Renders a detailed page for a specific artist, including their concert dates and locations.
	•	Error Handler: Handles errors and renders appropriate error pages.

### Templates

The HTML templates for the application are stored in the templates directory and include:

	•	index.html: Template for the home page.
	•	artist.html: Template for the artist details page.
	•	error.html: Template for error pages.

### Error Handling

The application handles various errors, such as invalid URLs, unsupported HTTP methods, and data fetching errors, by displaying appropriate error messages to the user.

Extensible and Scalable

The application is designed to be extensible, allowing for the easy addition of new features and improvements. The flexible nature of Go’s web framework enables further enhancements.

## Authors

	•	Ali Alqaed
	•	Sayed Ahmed Husain

## API Link

To see an example of the RESTful API used in this project, you can visit [Groupie Tracker API](https://groupietrackers.herokuapp.com/api).

## This project has helped me learn about:

	•	Manipulation and storage of data.
	•	JSON files and format.
	•	HTML.
	•	Event creation and display.
	•	Client-server communication.