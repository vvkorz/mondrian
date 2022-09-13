package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var tpl = template.Must(template.ParseFiles("form.html"))

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", imageHandler).Methods("GET")
	r.HandleFunc("/", imageHandler).Methods("POST")
	r.Handle("/img/mondrian_image.png", http.FileServer(http.Dir("./")))
	return r
}

func main() {
	r := newRouter()
	http.ListenAndServe(":8080", r)
}
