package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func About(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "About page")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Listening on port http://localhost%s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
