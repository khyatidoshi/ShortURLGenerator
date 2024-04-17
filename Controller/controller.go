package controller

import (
	"encoding/json"
	"net/http"

	model "github.com/khyatidoshi/ShortURLGenerator/Model"
	svc "github.com/khyatidoshi/ShortURLGenerator/Service"
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
	req := model.GenerateShortURLReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortURL, err := cnt.URLService.ShortenURL(req.LongURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"short_url": shortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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
