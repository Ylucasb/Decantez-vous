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
	case url[1] == "Sites" && len(url) == 3 && datamanagement.SitesExist(url[2]) && datamanagement.TestIsConnected(r):
		Sites(w, r, url[2])
	case url[1] == "InvalidPath" && len(url) == 2:
		InvalidPath(w, r)
	default:
		http.Redirect(w, r, "http://localhost:8080/InvalidPath", http.StatusSeeOther)
	}
}
