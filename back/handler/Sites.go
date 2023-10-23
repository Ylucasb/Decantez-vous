package handler

import (
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type structDisplaySites struct {
	IsNotValid bool
}

func Sites(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./front/html/sites.html"))

	structDisplaySites := structDisplaySites{}
	// allEmployees := datamanagement.RecuperationEmployee()
	// allWorplace := datamanagement.RecuperationWorkplace()
	// site := r.FormValue("selectWorkplace")

	err := t.Execute(w, structDisplaySites)
	if err != nil {
		log.Fatal(err)
	}
}
