package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

type structDisplayHome struct {
	AllTopics         []string
	HasAcceptedCookie string
	IsConnected       bool
	IsAdmin           bool
}

func getCookieValue(cookie *http.Cookie) string {
	var valueReturned string
	test := false
	value := fmt.Sprint(cookie)
	for _, element := range value {
		if test {
			valueReturned += string(element)
		}
		if element == 61 {
			test = true
		}
	}
	return valueReturned
}

func Home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./front/html/home.html", "./front/html/navBar.html"))
	structDisplayHome := structDisplayHome{}
	cookieConnected, _ := r.Cookie("isConnected")
	IsConnected := getCookieValue(cookieConnected)
	if IsConnected == "true" {
		structDisplayHome.IsConnected = true
	} else {
		structDisplayHome.IsConnected = false
	}
	t.ExecuteTemplate(w, "home", structDisplayHome)
}
