package handlers

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, "./templates/home.page.tmpl")
}

func About(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, "./templates/about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	temp, err := template.ParseFiles(tmpl, "./templates/base.layout.tmpl")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}
