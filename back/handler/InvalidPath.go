package handler

import (
	"net/http"
	"text/template"
)

type invalidPath struct {
	IsConnected bool
	IsAdmin     bool
}

func InvalidPath(w http.ResponseWriter, r *http.Request) {
	invalidPath := invalidPath{}
	t := template.Must(template.ParseFiles("./static/html/invalid.html"))
	t.ExecuteTemplate(w, "errorPath", invalidPath)
}
