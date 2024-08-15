package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port string = ":3000"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
}

type appConfig struct {
	useCache bool
}

func main() {

	app := application{}

	flag.BoolVar(&app.config.useCache, "cache", false, "Use template cache")
	flag.Parse()

	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Println("Starting app on port", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
