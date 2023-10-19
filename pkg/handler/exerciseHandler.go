package handler

import (
	"FitnessTracker/pkg/database"
	"log"
	"net/http"
	"strconv"
)

func CreateExercise(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		http.Redirect(w, r, "/admin/add", http.StatusSeeOther)
		return
	}

	muscleId, err := strconv.ParseInt(r.Form.Get("muscle"), 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		http.Redirect(w, r, "/admin/add", http.StatusSeeOther)
		return
	}

	e := database.Exercise{
		Name:        r.Form.Get("name"),
		Description: r.Form.Get("description"),
		MuscleId:    int(muscleId),
	}

	err = database.CreateExercise(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/admin/add", http.StatusSeeOther)
		return
	}

	log.Println("Exercise created successfully")
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
