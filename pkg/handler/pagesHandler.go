package handler

import (
	"log"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	renderPage(w, r, "./html/index.html")
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	renderPage(w, r, "./html/login.html")
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	renderPage(w, r, "./html/register.html")
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

func renderPage(w http.ResponseWriter, r *http.Request, page string) {
	tmpl, err := template.ParseFiles(page)
	if err != nil {
		NotFound(w, r)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
