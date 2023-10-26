package main

import "net/http"

func (app *application) route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheck)
	mux.HandleFunc("/v1/plants", app.getPlantsHandler)
	mux.HandleFunc("/v1/plant", app.createPlantHandler)
	mux.HandleFunc("/v1/plant/", app.deletePlantHandler)
	return mux
}
