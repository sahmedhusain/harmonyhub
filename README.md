# HarmonyHub 🎶

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![HTML](https://img.shields.io/badge/HTML-5-E34F26?style=flat&logo=html5)](https://developer.mozilla.org/en-US/docs/Web/HTML)
[![CSS](https://img.shields.io/badge/CSS-3-1572B6?style=flat&logo=css3)](https://developer.mozilla.org/en-US/docs/Web/CSS)
[![JavaScript](https://img.shields.io/badge/JavaScript-ES6+-F7DF1E?style=flat&logo=javascript)](https://developer.mozilla.org/en-US/docs/Web/JavaScript)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE.md)

Welcome to **HarmonyHub**, a sleek web application built with Go that brings together music lovers and concert enthusiasts. Discover detailed information about your favorite artists, their upcoming shows, and past performances—all in one place. Whether you're planning your next gig or just exploring new bands, HarmonyHub makes it easy and fun.

## ✨ Features

- **Artist Profiles** 📸: Dive deep into band details, including formation dates, first albums, and member lists.
- **Concert Tracking** 📅: Stay updated on concert dates and locations with real-time data.
- **Smart Search** 🔍: Find artists, members, locations, or dates instantly with our intuitive search bar featuring auto-suggestions.
- **User-Friendly Design** 💻: Clean, responsive interface that works on any device.
- **Reliable Performance** ⚡: Built to handle errors gracefully and keep running smoothly.

## 🛠️ Technologies Used

This project is built with a mix of powerful technologies:

- **Go** 🐹: Backend server, API handling, and data processing.
- **HTML** 🌐: Structure and layout of web pages.
- **CSS** 🎨: Styling for a beautiful, responsive design.
- **JavaScript** ⚙️: Interactive search functionality and dynamic suggestions.
- **REST API** 🔄: External data source for artist and concert information.

## 🎯 What We Aim For

HarmonyHub processes data from a dedicated API to create an engaging online experience. The API includes:

1. **Artists** 📸: Names, images, start years, debut albums, and band members.
2. **Locations** 📍: Where concerts happen.
3. **Dates** 📅: When the shows are scheduled.
4. **Relations** 🔗: How everything connects.

We use cards, tables, and interactive elements to visualize this data, ensuring smooth communication between the server and your browser.

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher installed on your machine.

### Installation

1. Clone the repo:
   ```bash
   git clone https://github.com/sahmedhusain/harmonyhub.git
   ```
2. Go to the project folder:
   ```bash
   cd harmonyhub
   ```
3. Run the app:

   ```bash
   go run main.go
   ```

4. Open your browser and head to `http://localhost:8080`.

## 📖 How to Use

Once running, you'll see the home page with a list of artists. Click on any artist for more details, or use the search bar to quickly find what you're looking for. The search supports names, members, places, and dates—it's case-insensitive and shows suggestions as you type!

### Search Examples

- Type "phil" and see suggestions like "Phil Collins - member" or "Phil Collins - artist/band".

## Screenshots

### Home Page

![Home Page](assets/img/home_screenshot.jpeg)
_Browse artists and use the search bar._

### Artist Details

![Artist Details](assets/img/SOJA-ArtistDetails.jpeg)
_Explore concert info and band history._

### Search Results

![Search Results](assets/img/Harmonysearch.jpeg)
_See live suggestions and results._

## 🛠️ Under the Hood

### Data Handling

We fetch and organize data from API endpoints, storing it in Go structs for quick access.

### Server Side

- **Handlers**: Manage home, artist pages, and errors.
- **Templates**: HTML files for rendering pages.
- **Error Management**: Keeps things stable with clear messages.

### Front End

- **Styling**: Follows best practices for consistency and ease of use.
- **JavaScript**: Powers the search with real-time suggestions.

The app is built to grow—easy to add new features thanks to Go's flexibility.

## 🔍 Search Bar Details

Our search tool lets you find specific info on the site:

- Searches: Artist/band names, members, locations, album dates, creation dates.
- Case-insensitive for easy typing.
- Shows suggestions as you write.
- Labels each suggestion (e.g., "Freddie Mercury - member").

## 🤝 Contributing

We'd love your help! Fork the repo, make changes, and send a pull request. Please follow Go standards and add tests where possible.

## 📄 License

Licensed under MIT - check [LICENSE.md](LICENSE.md) for more.

## 🙏 Acknowledgments

This project was created during a Go learning journey, emphasizing API work and web dev. Special thanks to the original API providers.

## 👥 Authors

- **Ali Alqaed**
- **Sayed Ahmed Husain** - [sayedahmed97.sad@gmail.com](mailto:sayedahmed97.sad@gmail.com)

## 🔗 API Reference

Check out the API we use: [Groupie Tracker API](https://groupietrackers.herokuapp.com/api).

## 📚 What I Learned

Building this taught me:

- How to handle and store data effectively.
- Working with JSON formats.
- Crafting HTML pages.
- Creating and showing events in web apps.
