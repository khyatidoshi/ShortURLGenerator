package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/khyati/ShortURLGenerator/Controller"
)

func main() {
	// Initialize router
	r := mux.NewRouter()

	handler := controller.NewURLController()
	// Define routes
	r.HandleFunc("/generate", handler.GenerateShortURLController).Methods("POST")
	// r.HandleFunc("/{short}", handler.RedirectController).Methods("GET")
	// r.HandleFunc("/stats/{short}", handler.GetStatsController()).Methods("GET")
	// r.HandleFunc("/delete/{short}", handler.DeleteShortURLController).Methods("DELETE")

	// Start HTTP server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
