package main

import (
	"FitnessTracker/pkg/auth"
	"FitnessTracker/pkg/database"
	"FitnessTracker/pkg/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.Index).Methods("GET")
	router.HandleFunc("/login", handler.LoginPage).Methods("GET")
	router.HandleFunc("/login", handler.Login).Methods("POST")
	router.HandleFunc("/register", handler.RegisterPage).Methods("GET")
	router.HandleFunc("/register", handler.Register).Methods("POST")
	router.HandleFunc("/home", auth.VerifyJWT(handler.HomePage)).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(handler.NotFound)

	err := database.InitDb()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println("Serving on port 8080")
	log.Println(http.ListenAndServe(":8080", router))
}
