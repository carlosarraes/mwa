package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, "./templates/home.page.tmpl")
}

func About(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, "./templates/about.page.tmpl")
}

func renderTemplateOld(w http.ResponseWriter, tmpl string) {
	temp, err := template.ParseFiles(tmpl, "./templates/base.layout.tmpl")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}

var tc = make(map[string]*template.Template)

func renderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error
	_, inMap := tc[t]

	if !inMap {
		fmt.Println("\nCreating template cache")
		err = createTemplateCache(t)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Using template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		t,
		"./templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		panic(err)
	}

	tc[t] = tmpl
	return nil
}
