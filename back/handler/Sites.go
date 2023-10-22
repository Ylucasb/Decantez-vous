package handler

import (
	"decantez-vous/back/datamanagement"
	"html/template"
	"log"
	"net/http"
	"time"
)

type structDisplaySites struct {
	IsNotValid bool
}

func Sites(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./front/html/sites.html"))
	structDisplaySites := structDisplaySites{}
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
			http.Redirect(w, r, "/sites", http.StatusSeeOther)
		}
	}
	cookieConnected, _ := r.Cookie("isConnected")
	IsConnected := datamanagement.GetCookieValue(cookieConnected)
	if IsConnected == "true" {
		structDisplaySites.IsNotValid = true
		// http.Redirect(w, r, "/acceuil", http.StatusSeeOther)
	} else {
		structDisplaySites.IsNotValid = false

	}
	err := t.Execute(w, structDisplaySites)
	if err != nil {
		log.Fatal(err)
	}
}
