package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/khyatidoshi/ShourtURLGenerator/Controller"
)

func main() {
	// Initialize router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/generate", Controller.ShortURLController()()).Methods("POST")
	r.HandleFunc("/{short}", Controller.RedirectController()).Methods("GET")
	r.HandleFunc("/stats/{short}", Controller.StatsController()).Methods("GET")
	r.HandleFunc("/delete/{short}", Controller.DeleteShortURLController).Methods("DELETE")

	// Start HTTP server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
