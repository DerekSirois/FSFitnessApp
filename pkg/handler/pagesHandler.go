package handler

import (
	"FitnessTracker/pkg/database"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type AdminExercise struct {
	Id          int
	Name        string
	Description string
	Muscle      string
}

type EditExercise struct {
	Muscle   []*database.Muscle
	Exercise *database.Exercise
}

func Index(w http.ResponseWriter, r *http.Request) {
	renderPage(w, r, "./html/index.html", nil)
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	renderPage(w, r, "./html/login.html", nil)
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	renderPage(w, r, "./html/register.html", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	renderPage(w, r, "./html/home.html", nil)
}

func AdminPage(w http.ResponseWriter, r *http.Request) {
	e, err := database.GetAllExercises()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	ae, err := MapExercise(e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	renderPage(w, r, "./html/admin/admin.html", ae)
}

func EditExercisePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		http.Redirect(w, r, "/admin/edit", http.StatusSeeOther)
		return
	}

	m, err := database.GetAllMuscle()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	e, err := database.GetByIdExercise(int(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	ee := &EditExercise{Muscle: m, Exercise: e}

	renderPage(w, r, "./html/admin/editExercise.html", ee)
}

func AddExercisePage(w http.ResponseWriter, r *http.Request) {
	b, err := database.GetAllMuscle()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	renderPage(w, r, "./html/admin/addExercise.html", b)
}

func NotFound(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("./html/404.html")
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func renderPage(w http.ResponseWriter, r *http.Request, page string, data any) {
	tmpl, err := template.ParseFiles(page)
	if err != nil {
		NotFound(w, r)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func MapExercise(e []*database.Exercise) ([]*AdminExercise, error) {
	ae := make([]*AdminExercise, 0)
	for _, v := range e {
		muscle, err := database.GetByIdMuscle(v.MuscleId)
		if err != nil {
			return nil, err
		}
		ae = append(ae, &AdminExercise{
			Id:          v.Id,
			Name:        v.Name,
			Description: v.Description,
			Muscle:      muscle.Name,
		})
	}
	return ae, nil
}
