package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/spargonaut/plant_info_service/internal/data"
	"io"
	"net/http"
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

	plants, err := app.profiles.Plants.GetAll()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
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
		Name                  string  `json:"name,omitempty" validate:"required"`
		CommonName            string  `json:"common_name" validate:"required"`
		SeedCompany           string  `json:"seed_company" validate:"required"`
		ExpectedDaysToHarvest int32   `json:"expected_days_to_harvest" validate:"required"`
		Type                  string  `json:"type" validate:"oneof=harvest_once cut_and_come_again"`
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

	err = validate.Struct(input)
	if err != nil {
		fmt.Println("there was a validation error")
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("field %s is %s\n", err.Field(), err.ActualTag())
			if err.Value() != "" {
				fmt.Printf("\"%s\" not found in %s ", err.Value(), err.Param())
			}
			fmt.Println()
		}

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	plant := &data.Plant{
		Name:                  input.Name,
		CommonName:            input.CommonName,
		SeedCompany:           input.SeedCompany,
		ExpectedDaysToHarvest: input.ExpectedDaysToHarvest,
		Type:                  input.Type,
		PhLow:                 input.PhLow,
		PhHigh:                input.PhHigh,
		ECLow:                 input.ECLow,
		ECHigh:                input.ECHigh,
	}

	err = app.profiles.Plants.Insert(plant)
	if err != nil {
		fmt.Println("Insert Error")
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("plant profile created\n"))
}
