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

	fmt.Printf("Listening on port http://localhost%s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
