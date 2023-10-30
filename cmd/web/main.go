package main

import (
	"flag"
	"github.com/spargonaut/plant_info_service/internal/models"
	"log"
	"net/http"
)

type application struct {
	plantInfo *models.PlantProfile
}

func main() {
	addr := flag.String("addr", ":8090", "HTTP network address")
	commandEndpoint := flag.String("commandEndpoint", "http://localhost:4000/v1/plant", "CommandEndpoint for the Plant Info command service")
	flag.Parse()

	app := &application{
		plantInfo: &models.PlantProfile{CommandEndpoint: *commandEndpoint},
	}
	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Printf("Starting the server on %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
