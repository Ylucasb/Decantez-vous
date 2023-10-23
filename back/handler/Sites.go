package handler

import (
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type structDisplaySites struct {
	IsNotValid bool
}


func Sites(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./front/html/sites.html"))

	allSites := datamanagement.sitesPageStruct{}
	allEmployees := datamanagement.recuperationEmployee()
	allWorplace := datamanagement.recuperationWorkplace()
	site := r.FormValue("selectWorkplace")
	
}
