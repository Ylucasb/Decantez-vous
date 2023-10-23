package handler

import (
	"decantez-vous/back/datamanagement"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type structDisplaySites struct {
	Employee []datamanagement.EmployeeFromDb
}

func Sites(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./front/html/sites.html"))

	structDisplaySites := structDisplaySites{datamanagement.RecuperationEmployee()}
	// allWorplace := datamanagement.RecuperationWorkplace()
	// site := r.FormValue("selectWorkplace")

	err := t.Execute(w, structDisplaySites)
	if err != nil {
		log.Fatal(err)
	}
}
