function searchSuggestions() {
    const query = document.getElementById('searchBar').value;
    const suggestions = document.getElementById('suggestions');

    if (query.trim() === '') {
        suggestions.style.display = 'none'; // Hide suggestions if query is empty
        return;
    }

    fetch(`/search?q=${encodeURIComponent(query)}`)
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            suggestions.innerHTML = '';
            if (data.length > 0) {
                data.forEach(item => {
                    let div = document.createElement('div');
                    div.innerHTML = `${item.name} - ${item.type}`;
                    div.dataset.id = item.id; // Store the artist ID in a data attribute
                    div.classList.add('suggestion');
                    div.addEventListener('click', function () {
                        window.location.href = `/artist?id=${item.id}`; // Redirect to the artist page
                    });
                    suggestions.appendChild(div);
                });
                suggestions.style.display = 'block'; // Show suggestions
            } else {
                suggestions.style.display = 'none'; // Hide suggestions if no data
            }
        })
        .catch(error => {
            console.error('There was a problem with the fetch operation:', error);
            suggestions.style.display = 'none'; // Hide suggestions on error
        });
}

// Handle click event on suggestions
document.addEventListener('click', function (event) {
    const searchBar = document.getElementById('searchBar');
    const suggestions = document.getElementById('suggestions');

    if (!searchBar.contains(event.target) && !suggestions.contains(event.target)) {
        suggestions.style.display = 'none';
    }
});
