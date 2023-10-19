package handler

import (
	"decantez-vous/back/datamanagement"
	"html/template"
	"net/http"
	"time"
)

type structDisplayHome struct {
	IsNotValid bool
}

func Home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./front/html/home.html", "./front/html/navBar.html"))
	structDisplayHome := structDisplayHome{}
	email := r.FormValue("userInput")
	userPassword := r.FormValue("userPassword")
	if email != "" && userPassword != "" {
		ifUserExist, idUser := datamanagement.IsRegister(email, userPassword)
		if ifUserExist {
			cookieIdUser := http.Cookie{Name: "idUser", Value: idUser, Expires: time.Now().Add(1 / 2 * time.Hour)}
			http.SetCookie(w, &cookieIdUser)
			cookieIsConnected := http.Cookie{Name: "isConnected", Value: "true", Expires: time.Now().Add(1 / 2 * time.Hour)}
			http.SetCookie(w, &cookieIsConnected)
		}
	}
	cookieConnected, _ := r.Cookie("isConnected")
	IsConnected := datamanagement.GetCookieValue(cookieConnected)
	if IsConnected == "true" {
		structDisplayHome.IsNotValid = true
	} else {
		structDisplayHome.IsNotValid = false
	}
	t.ExecuteTemplate(w, "home", structDisplayHome)
}
