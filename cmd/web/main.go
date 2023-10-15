package main

import (
	"FitnessTracker/pkg/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.Index).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(handler.NotFound)
	log.Println("Serving on port 8080")
	log.Println(http.ListenAndServe(":8080", router))
}
