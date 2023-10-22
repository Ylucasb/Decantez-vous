package main

import (
	"decantez-vous/back/handler"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var port = ":8080"

func main() {
	//handlers
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/acceuil", handler.Acceuil)
	http.HandleFunc("/sites", handler.Sites)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("front/css"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("front/fonts"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("front/img"))))
	fmt.Println("(http://localhost"+port+"/"+") - Server started on port", port)
	http.ListenAndServe(port, nil)
}
