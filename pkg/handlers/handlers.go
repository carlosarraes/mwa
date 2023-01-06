package handlers

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, "./templates/home.html")
}

func About(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, "./templates/about.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	temp, err := template.ParseFiles(tmpl)
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}
