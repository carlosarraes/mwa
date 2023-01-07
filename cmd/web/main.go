package main

import (
	"fmt"
	"mwa/pkg/config"
	"mwa/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := handlers.CreateTemplateCache()
	if err != nil {
		fmt.Println("Something went wrong")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	handlers.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Listening on port http://localhost%s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
