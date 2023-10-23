package handler

import (
	"decantez-vous/back/datamanagement"
	"net/http"
	"strings"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.String(), "/")
	switch true {
	case url[1] == "" && len(url) == 2:
		Connection(w, r)
	case url[1] == "Sites" && len(url) == 2 && datamanagement.TestIsRegister(r):
		Sites(w, r)
	default:
		if len(url) >= 3 {
			http.Redirect(w, r, "http://localhost:8080/InvalidPath", http.StatusSeeOther)
		}
		InvalidPath(w, r)
	}
}
