package handler

import (
	"decantez-vous/back/datamanagement"
	"html/template"
	"log"
	"net/http"
	"time"
)

type structDisplayHome struct {
	IsNotValid bool
}

func Connection(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./front/html/home.html"))
	structDisplayHome := structDisplayHome{}
	structDisplayHome.IsNotValid = false
	email := r.FormValue("userEmail")
	userPassword := r.FormValue("userPassword")
	if email != "" && userPassword != "" {
		ifUserExist, idUser := datamanagement.IsRegister(email, userPassword)
		if ifUserExist {
			cookieIdUser := http.Cookie{Name: "idUser", Value: idUser, Expires: time.Now().Add(30 * time.Minute)}
			http.SetCookie(w, &cookieIdUser)
			cookieIsConnected := http.Cookie{Name: "isConnected", Value: "true", Expires: time.Now().Add(30 * time.Minute)}
			http.SetCookie(w, &cookieIsConnected)
			http.Redirect(w, r, "/acceuil", http.StatusSeeOther)
		} else {
			structDisplayHome.IsNotValid = true
		}
	}
	err := t.Execute(w, structDisplayHome)
	if err != nil {
		log.Fatal(err)
	}
}
