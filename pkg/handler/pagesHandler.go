package handler

import (
	"FitnessTracker/pkg/database"
	"log"
	"net/http"
	"text/template"
)

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
	renderPage(w, r, "./html/admin/admin.html", nil)
}

func AddExercisePage(w http.ResponseWriter, r *http.Request) {
	b, err := database.GetAllBodyPart()
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
