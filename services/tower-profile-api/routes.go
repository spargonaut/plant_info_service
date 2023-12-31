package main

import "net/http"

func (app *application) route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheck)
	mux.HandleFunc("/v1/towers", app.getGrowTowersHandler)
	mux.HandleFunc("/v1/tower", app.createGrowTowerHandler)
	mux.HandleFunc("/v1/tower/", app.deleteGrowTowerHandler)
	return mux
}
