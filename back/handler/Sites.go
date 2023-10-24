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
	workplace := r.FormValue("selectWorkplace")
	print(workplace)

	// structDisplaySites := structDisplaySites{datamanagement.RecuperationEmployee()}

	structDisplaySites := structDisplaySites{datamanagement.RecuperationEmployeeWorkplace(workplace)}
	// allWorplace := datamanagement.RecuperationWorkplace()

	err := t.Execute(w, structDisplaySites)
	if err != nil {
		log.Fatal(err)
	}
}
