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

func Home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./front/html/home.html"))
	structDisplayHome := structDisplayHome{}
	email := r.FormValue("userEmail")
	userPassword := r.FormValue("userPassword")
	if email != "" && userPassword != "" {
		ifUserExist, idUser := datamanagement.IsRegister(email, userPassword)
		if ifUserExist {
			cookieIdUser := http.Cookie{Name: "idUser", Value: idUser, Expires: time.Now().Add(1 / 2 * time.Hour)}
			http.SetCookie(w, &cookieIdUser)
			cookieIsConnected := http.Cookie{Name: "isConnected", Value: "true", Expires: time.Now().Add(1 / 2 * time.Hour)}
			http.SetCookie(w, &cookieIsConnected)
		} else {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
		}
	}
	cookieConnected, _ := r.Cookie("isConnected")
	IsConnected := datamanagement.GetCookieValue(cookieConnected)
	if IsConnected == "true" {
		structDisplayHome.IsNotValid = true
		// http.Redirect(w, r, "/acceuil", http.StatusSeeOther)
	} else {
		structDisplayHome.IsNotValid = false

	}
	err := t.Execute(w, structDisplayHome)
	if err != nil {
		log.Fatal(err)
	}
}
