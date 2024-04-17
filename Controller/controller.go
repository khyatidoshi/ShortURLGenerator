package comtroller

import (
	"fmt"
	"net/http"
)

// Controllers
func ShortURLController(w http.ResponseWriter, r *http.Request) {
	fmt.Printf(" Something started")
	// Generate short URL logic
}

func RedirectController(w http.ResponseWriter, r *http.Request) {
	// Redirect short URL logic
}

func StatsController(w http.ResponseWriter, r *http.Request) {
	// Access counts logic
}

func DeleteShortURLController(w http.ResponseWriter, r *http.Request) {
	// Delete short URL logic
}
