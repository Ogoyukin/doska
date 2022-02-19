package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", Handler)
	http.ListenAndServe(":8090", nil)
}

var counter = 0

type Params struct {
	Counter int
	Name    string
}

func Handler(w http.ResponseWriter, r *http.Request) {

	counter++
	tmp, _ := template.ParseFiles("templates/index.html")
	tmp.ParseGlob("templates/*.html")

	name := r.URL.Query().Get("name")

	tmp.Execute(w, Params{
		Counter: counter,
		Name:    name,
	})
}
