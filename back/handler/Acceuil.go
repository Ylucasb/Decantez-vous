package handler

import (
	"html/template"
	"log"
	"net/http"
)

func Acceuil(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./front/html/acceuil.html"))

	err := t.Execute(w, r)
	if err != nil {
		log.Fatal(err)
	}
}
