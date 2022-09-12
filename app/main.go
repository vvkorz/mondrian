package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("form.html"))

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")
	r.HandleFunc("/", imageHandler).Methods("POST")
	r.Handle("/img/mondrian_image.png", http.FileServer(http.Dir("./")))
	return r
}

func main() {
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}
