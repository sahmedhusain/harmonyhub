function searchSuggestions() {
    const query = document.getElementById('searchBar').value;
    // Gets the current value of the search bar
    const suggestions = document.getElementById('suggestions');
    // Gets the suggestions container

    if (query.trim() === '') {
        suggestions.style.display = 'none';
        // Hides suggestions if query is empty
        return;
    }

    fetch(`/search?q=${encodeURIComponent(query)}`)
        // Fetches search results from the server
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
                // Throws an error if the response is not ok
            }
            return response.json();
            // Returns the response in JSON format
        })
        .then(data => {
            suggestions.innerHTML = '';
            // Clears previous suggestions
            if (data.length > 0) {
                data.forEach(item => {
                    let div = document.createElement('div');
                    // Creates a new div for each suggestion
                    div.innerHTML = `${item.name} - ${item.type}`;
                    // Sets the inner HTML of the suggestion
                    div.dataset.id = item.id;
                    // Stores the artist ID in a data attribute
                    div.classList.add('suggestion');
                    // Adds a class to the suggestion
                    div.addEventListener('click', function () {
                        window.location.href = `/artist?id=${item.id}`;
                        // Redirects to the artist page on click
                    });
                    suggestions.appendChild(div);
                    // Appends the suggestion to the suggestions container
                });
                suggestions.style.display = 'block';
                // Shows the suggestions
            } else {
                suggestions.style.display = 'none';
                // Hides suggestions if no data
            }
        })
        .catch(error => {
            console.error('There was a problem with the fetch operation:', error);
            // Logs any fetch errors to the console
            suggestions.style.display = 'none';
            // Hides suggestions on error
        });
}

// Handle click event on suggestions
document.addEventListener('click', function (event) {
    const searchBar = document.getElementById('searchBar');
    const suggestions = document.getElementById('suggestions');

    if (!searchBar.contains(event.target) && !suggestions.contains(event.target)) {
        suggestions.style.display = 'none';
        // Hides suggestions if click is outside search bar and suggestions
    }
});
