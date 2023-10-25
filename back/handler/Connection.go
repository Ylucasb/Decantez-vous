package handler

import (
	"decantez-vous/back/datamanagement"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type structDisplayHome struct {
	IsNotValid bool
}

func getWorkPlace(idUser int) string {
	rows := datamanagement.SelectDB("SELECT workplace.adress FROM workplace JOIN employee ON workplace.idWorkplace = employee.idWorkplace WHERE employee.idEmployee=?", int(idUser))
	var adress string
	for rows.Next() {
		rows.Scan(&adress)
	}
	return adress
}

func Connection(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./front/html/home.html"))
	structDisplayHome := structDisplayHome{}
	structDisplayHome.IsNotValid = false
	email := r.FormValue("userEmail")
	userPassword := r.FormValue("userPassword")
	if email != "" && userPassword != "" {
		ifUserExist, idUser, isPays := datamanagement.IsRegister(email, userPassword)
		if ifUserExist {
			cookieIdUser := http.Cookie{Name: "idUser", Value: strconv.Itoa(idUser), Expires: time.Now().Add(30 * time.Minute)}
			http.SetCookie(w, &cookieIdUser)
			cookieIsConnected := http.Cookie{Name: "isConnected", Value: "true", Expires: time.Now().Add(30 * time.Minute)}
			http.SetCookie(w, &cookieIsConnected)
			cookieIsPays := http.Cookie{Name: "isPays", Value: strconv.FormatBool(isPays), Expires: time.Now().Add(30 * time.Minute)}
			http.SetCookie(w, &cookieIsPays)
			http.Redirect(w, r, "/Sites/"+getWorkPlace(idUser), http.StatusSeeOther)
		} else {
			structDisplayHome.IsNotValid = true
		}
	}
	err := t.Execute(w, structDisplayHome)
	if err != nil {
		log.Fatal(err)
	}
}
