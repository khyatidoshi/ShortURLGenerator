package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	controller "github.com/khyatidoshi/ShortURLGenerator/server/Controller"
)

func main() {
	// Initialize router
	r := mux.NewRouter()

	handler := controller.NewURLController()
	// Define routes
	r.HandleFunc("/generate", handler.GenerateShortURLController).Methods("POST")
	r.HandleFunc("/{short}", handler.RedirectController).Methods("GET")
	r.HandleFunc("/stats/{short}", handler.GetStatsController).Methods("GET")

	go startScheduledDeleteTasks(handler)
	go startKafkaConsumer(handler)

	// Start HTTP server
	log.Println("Server started on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}

func startScheduledDeleteTasks(handler *controller.URLController) {
	fmt.Println("deleting expired URLs scheduled at : ", time.Now())
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := handler.URLService.DeleteURL()
			if err != nil {
				log.Printf("Error during scheduled deletion of expired URLs: %v", err)
			} else {
				log.Println("Expired URLs successfully deleted.")
			}
		}
	}
}

func startKafkaConsumer(handler *controller.URLController) {
	fmt.Println("started Kafka consumer at : ", time.Now())
	ticker := time.NewTicker(15 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			handler.URLService.ConsumeMessage()
		}
	}
}
