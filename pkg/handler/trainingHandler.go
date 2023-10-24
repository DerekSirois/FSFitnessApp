package handler

import (
	"FitnessTracker/pkg/auth"
	"FitnessTracker/pkg/database"
	"log"
	"net/http"
)

func CreateTraining(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		http.Redirect(w, r, "/training", http.StatusSeeOther)
		return
	}

	userId, err := auth.GetAuthenticatedUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/training", http.StatusSeeOther)
		return
	}

	t := database.Training{
		Name:    r.Form.Get("name"),
		Weekday: r.Form.Get("weekday"),
		UserId:  userId,
	}

	err = database.CreateTraining(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/training", http.StatusSeeOther)
		return
	}

	log.Println("Training created successfully")
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
