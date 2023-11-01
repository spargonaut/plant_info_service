package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/plant/create", app.createPlant)
	mux.HandleFunc("/plant/delete", app.deletePlant)

	return mux
}
