package handler

import (
	"FitnessTracker/pkg/auth"
	"FitnessTracker/pkg/database"
	"FitnessTracker/pkg/utils"
	"log"
	"net/http"
)

type Home struct {
	Token string
}

func Register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	hash, err := auth.HashPassword(r.Form.Get("password"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	u := database.User{
		Username: r.Form.Get("username"),
		Email:    r.Form.Get("email"),
		Password: hash,
	}

	err = u.Create()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	log.Println("Registered successfully")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	u := database.User{}
	err = u.GetByUsername(r.Form.Get("username"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if !auth.CheckPasswordHash(r.Form.Get("password"), u.Password) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	token, err := auth.CreateJWTToken(u.Id, u.Username)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	session, _ := utils.Store.Get(r, "session")
	session.Values["token"] = token

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
