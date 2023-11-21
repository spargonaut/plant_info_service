package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/spargonaut/plant_info_service/internal/data"
	"io"
	"net/http"
)

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

type envelope map[string]any

func (app *application) getGrowTowersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	towers, err := app.profiles.GrowTowers.GetAll()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	js, err := json.MarshalIndent(envelope{"towers": towers}, "", "\t")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		fmt.Println("Error writing the GET growtowers response.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (app *application) createGrowTowerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Name         string  `json:"name" validate:"required"`
		Type         string  `json:"type"  validate:"required"`
		TargetPhLow  float32 `json:"target_ph_low,omitempty"`
		TargetPhHigh float32 `json:"target_ph_high,omitempty"`
		TargetECLow  float32 `json:"target_ec_low,omitempty"`
		TargetECHigh float32 `json:"target_ec_high,omitempty"`
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

	tower := &data.GrowTower{
		Name:         input.Name,
		Type:         input.Type,
		TargetPhLow:  input.TargetPhLow,
		TargetPhHigh: input.TargetPhHigh,
		TargetECLow:  input.TargetECLow,
		TargetECHigh: input.TargetECHigh,
	}

	err = app.profiles.GrowTowers.Insert(tower)
	if err != nil {
		fmt.Println("Insert Error")
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("grow tower profile created\n"))
}
