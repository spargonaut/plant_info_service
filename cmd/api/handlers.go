package main

import (
	"encoding/json"
	"fmt"
	"io"
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

	var input struct {
		Name                  string  `json:"name"`
		CommonName            string  `json:"common_name"`
		SeedCompany           string  `json:"seed_company"`
		ExpectedDaysToHarvest int32   `json:"expected_days_to_harvest"`
		Type                  string  `json:"(harvest once|cut and come again)"`
		PhLow                 float32 `json:"ph_low"`
		PhHigh                float32 `json:"ph_high"`
		ECLow                 float32 `json:"ec_low"`
		ECHigh                float32 `json:"ec_high"`
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Read Error")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &input)
	if err != nil {
		fmt.Printf("Unmarshall Error")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%v\n", input)
}
