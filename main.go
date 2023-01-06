package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

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

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Listening on port http://localhost%s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
