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
	router.HandleFunc("/home", auth.VerifyJWT(handler.HomePage, false)).Methods("GET")
	router.HandleFunc("/admin", auth.VerifyJWT(handler.AdminPage, true)).Methods("GET")
	router.HandleFunc("/admin/add", auth.VerifyJWT(handler.AddExercisePage, true)).Methods("GET")
	router.HandleFunc("/admin/add", auth.VerifyJWT(handler.CreateExercise, true)).Methods("POST")
	router.HandleFunc("/admin/edit/{id:[0-9]+}", auth.VerifyJWT(handler.EditExercisePage, true)).Methods("GET")
	router.HandleFunc("/admin/edit/{id:[0-9]+}", auth.VerifyJWT(handler.UpdateExercise, true)).Methods("POST")
	router.HandleFunc("/admin/{id:[0-9]+}", auth.VerifyJWT(handler.DeleteExercise, true)).Methods("POST")
	router.HandleFunc("/training", auth.VerifyJWT(handler.AddTrainingPage, false)).Methods("GET")
	router.HandleFunc("/training", auth.VerifyJWT(handler.CreateTraining, false)).Methods("POST")
	router.HandleFunc("/training/delete/{id:[0-9]+}", auth.VerifyJWT(handler.DeleteTraining, false)).Methods("POST")

	router.NotFoundHandler = http.HandlerFunc(handler.NotFound)

	err := database.InitDb()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println("Serving on port 8080")
	log.Println(http.ListenAndServe(":8080", router))
}
