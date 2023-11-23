package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/towers", app.towerHome)
	mux.HandleFunc("/plant/create", app.createPlant)
	mux.HandleFunc("/plant/delete", app.deletePlant)

	return mux
}
