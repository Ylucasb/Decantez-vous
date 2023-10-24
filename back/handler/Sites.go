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
	IsPays   bool
}

func Sites(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./front/html/sites.html"))
	cookieIsPays, _ := r.Cookie("isPays")
	isPays := datamanagement.GetCookieValue(cookieIsPays)
	var isPaysBool bool
	if isPays == "true" {
		isPaysBool = true
	} else {
		isPaysBool = false
	}
	structDisplaySites := structDisplaySites{datamanagement.RecuperationEmployee(), isPaysBool}
	if isPaysBool {
		for i := 0; i < len(structDisplaySites.Employee); i++ {
			structDisplaySites.Employee[i].IsPays = isPaysBool
		}
	}
	// allWorplace := datamanagement.RecuperationWorkplace()
	// site := r.FormValue("selectWorkplace")

	err := t.Execute(w, structDisplaySites)
	if err != nil {
		log.Fatal(err)
	}
}
