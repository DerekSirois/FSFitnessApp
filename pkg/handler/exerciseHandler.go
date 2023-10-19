package handler

import (
	"FitnessTracker/pkg/database"
	"github.com/gorilla/mux"
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

func UpdateExercise(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		http.Redirect(w, r, "/admin/edit", http.StatusSeeOther)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/admin/edit", http.StatusSeeOther)
		return
	}

	muscleId, err := strconv.ParseInt(r.Form.Get("muscle"), 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		http.Redirect(w, r, "/admin/add", http.StatusSeeOther)
		return
	}

	e := &database.Exercise{
		Id:          int(id),
		Name:        r.Form.Get("name"),
		Description: r.Form.Get("description"),
		MuscleId:    int(muscleId),
	}

	err = database.UpdateExercise(e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/admin/edit", http.StatusSeeOther)
		return
	}

	log.Println("Exercise updated successfully")
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func DeleteExercise(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	err = database.DeleteExercise(int(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	log.Println("Exercise deleted successfully")
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
