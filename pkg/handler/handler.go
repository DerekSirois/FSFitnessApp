package handler

import (
	"log"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./html/index.html")
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