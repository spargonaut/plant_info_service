package main

import (
	"flag"
	"github.com/spargonaut/plant_info_service/internal/models"
	"log"
	"net/http"
)

type application struct {
	plantList *models.PlantListModel
}

func main() {
	addr := flag.String("addr", ":8090", "HTTP network address")
	endpoint := flag.String("endpoint", "http://localhost:4000/v1/plant", "Endpoint for the Plant Info command service")
	flag.Parse()

	app := &application{
		plantList: &models.PlantListModel{Endpoint: *endpoint},
	}
	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Printf("Starting the server on %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
