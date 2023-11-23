package main

import (
	"flag"
	"github.com/spargonaut/plant_info_service/internal/models"
	"log"
	"net/http"
)

type application struct {
	plantInfo *models.PlantProfile
	towerInfo *models.GrowTowerProfile
}

func main() {
	addr := flag.String("addr", ":8090", "HTTP network address")
	plantCmd := flag.String("plantCmd", "http://localhost:4000/v1/plant", "Command Endpoint for the Plant Info command service")
	plantQry := flag.String("plantQry", "http://localhost:4000/v1/plants", "Query Endpoint for the Plant Info command service")
	towerCmd := flag.String("towerCmd", "http://localhost:4001/v1/tower", "Command Endpoint for the Grow Tower Info command service")
	towerQry := flag.String("towerQry", "http://localhost:4001/v1/towers", "Query Endpoint for the Grow Tower Info command service")
	flag.Parse()

	app := &application{
		plantInfo: &models.PlantProfile{
			CommandEndpoint: *plantCmd,
			QueryEndpoint:   *plantQry,
		},
		towerInfo: &models.GrowTowerProfile{
			CommandEndpoint: *towerCmd,
			QueryEndpoint:   *towerQry,
		},
	}
	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Printf("Starting the server on %s", *addr)
	log.Printf("plant command endpoint: %s", app.plantInfo.CommandEndpoint)
	log.Printf("plant query endpoint: %s", app.plantInfo.QueryEndpoint)
	log.Printf("tower command endpoint: %s", app.towerInfo.CommandEndpoint)
	log.Printf("tower query endpoint: %s", app.towerInfo.QueryEndpoint)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
