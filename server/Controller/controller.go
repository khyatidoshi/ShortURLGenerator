package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	model "github.com/khyati/ShortURLGenerator/server/Model"
	svc "github.com/khyati/ShortURLGenerator/server/Service"
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

func (cnt *URLController) RedirectController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	longURL, err := cnt.URLService.GetLongURL(vars["shortUrl"])
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, longURL, http.StatusMovedPermanently)
}

func StatsController(w http.ResponseWriter, r *http.Request) {
	// Access counts logic
}

func (cnt *URLController) DeleteShortURLController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := cnt.URLService.DeleteURL(vars["shortUrl"])
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
