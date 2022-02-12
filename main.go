package main

import (
	"github.com/Ogoyukin/doska/app"
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
	config := app.Config{
		Name: "iPhone 13 Pro",
	}

	counter++
	tmp, _ := template.ParseFiles("templates/index.html")

	name := r.URL.Query().Get("name")
	if name == "" {
		name = config.Name
	}

	tmp.Execute(w, Params{
		Counter: counter,
		Name:    name,
	})
}
