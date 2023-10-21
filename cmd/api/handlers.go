package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/spargonaut/plant_info_service/internal/data"
)

type envelope map[string]any

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	healthData := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	js, err := json.Marshal(healthData)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		fmt.Println("Error writing the healthcheck response.")
		return
	}
}

func (app *application) getPlantsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	plants := []data.Plant{
		{
			ID:                    23,
			CreatedAt:             time.Now(),
			Name:                  "Speedy - Salad Arugula - Gourmet Greens",
			CommonName:            "Speedy Arugula",
			SeedCompany:           "Territorial",
			ExpectedDaysToHarvest: 30,
			Type:                  "harvest once",
			PhLow:                 6,
			PhHigh:                7.5,
			ECLow:                 0.8,
			ECHigh:                1.2,
			Version:               1,
		},
	}

	js, err := json.MarshalIndent(envelope{"plants": plants}, "", "\t")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		fmt.Println("Error writing the GET plants response.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (app *application) createPlantHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "Added a new plant")
}
