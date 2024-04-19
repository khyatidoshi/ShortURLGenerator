package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	model "github.com/khyatidoshi/ShortURLGenerator/server/Model"
	svc "github.com/khyatidoshi/ShortURLGenerator/server/Service"
	utils "github.com/khyatidoshi/ShortURLGenerator/server/Utils"
)

// Controllers

type URLController struct {
	URLService *svc.URLService
}

func NewURLController() *URLController {
	return &URLController{
		URLService: svc.NewURLService(),
	}
}
func (cnt *URLController) GenerateShortURLController(w http.ResponseWriter, r *http.Request) {
	req := model.GenerateShortURLReq{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate URL
	if err := utils.ValidateURL(req.LongURL); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	expiryDate, err := utils.ValidateExpiryDate(req.ExpiryDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortURL, err := cnt.URLService.ShortenURL(req, expiryDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"short_url": shortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (cnt *URLController) RedirectController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["short"]

	// Check if the short URL is empty
	if shortURL == "" {
		http.Error(w, "Short URL is empty", http.StatusBadRequest)
		return
	}

	longURL, err := cnt.URLService.GetLongURL(shortURL)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, longURL, http.StatusMovedPermanently)
}

func StatsController(w http.ResponseWriter, r *http.Request) {
	// Access counts logic
}
