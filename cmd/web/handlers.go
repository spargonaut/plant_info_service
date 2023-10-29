package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "the home page")
}

func (app *application) createPlant(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.createPlantForm(w, r)
	case http.MethodPost:
		app.createPlantProcess(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (app *application) createPlantForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><head><title>Create Plant</title></head>"+
		"<body><h1>Create Plant</h1><form action=\"/plant/create\" method=\"post\">"+
		"<label for=\"name\">Name</label><input type=\"text\" name=\"name\" id=\"name\">"+
		"<label for=\"common_name\">Common Name</label><input type=\"text\" name=\"common_name\" id=\"common_name\">"+
		"<label for=\"seed_company\">Seed Company</label><input type=\"text\" name=\"seed_company\" id=\"seed_company\">"+
		"<label for=\"expected_days_to_harvest\">Expected Days To Harvests</label><input type=\"number\" step=\"1\" name=\"expected_days_to_harvest\" id=\"expected_days_to_harvest\">"+
		"<label for=\"type\">Type</label><input type=\"text\" name=\"type\" id=\"type\">"+
		"<label for=\"ph_low\">PH Low</label><input type=\"number\" step=\"0.1\" name=\"ph_low\" id=\"ph_low\">"+
		"<label for=\"ph_high\">PH High</label><input type=\"number\" step=\"0.1\" name=\"ph_high\" id=\"ph_high\">"+
		"<label for=\"ec_low\">EC Low</label><input type=\"number\" step=\"0.1\" name=\"ec_low\" id=\"ec_low\">"+
		"<label for=\"ec_high\">EC High</label><input type=\"number\" step=\"0.1\" name=\"ec_high\" id=\"ec_high\">"+
		"<button type=\"submit\">Create</button></form></body></html>")
}

func (app *application) createPlantProcess(w http.ResponseWriter, r *http.Request) {
	fmt.Println("processing the create plant form")
	name := r.PostFormValue("name")
	if name == "" {
		fmt.Println("name error")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	common_name := r.PostFormValue("common_name")
	if common_name == "" {
		fmt.Println("common_name error")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	seed_company := r.PostFormValue("seed_company")
	if seed_company == "" {
		fmt.Println("seed company")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	expected_days_to_harvest, err := strconv.Atoi(r.PostFormValue("expected_days_to_harvest"))
	if expected_days_to_harvest < 1 || err != nil {
		fmt.Println("expected days error")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	harvestType := r.PostFormValue("type")
	if harvestType != "harvest_once" && harvestType != "cut_and_come_again" {
		fmt.Println("harvest type error")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	phLowFloat, err := strconv.ParseFloat(r.PostFormValue("ph_low"), 32)
	if phLowFloat < 0.1 || err != nil {
		fmt.Println("ph low error")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	phLow := float32(phLowFloat)

	phHighFloat, err := strconv.ParseFloat(r.PostFormValue("ph_high"), 32)
	if phHighFloat < 0.1 || err != nil {
		fmt.Println("ph_high error")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	phHigh := float32(phHighFloat)

	ecLowFloat, err := strconv.ParseFloat(r.PostFormValue("ec_low"), 32)
	if ecLowFloat < 0.1 || err != nil {
		fmt.Println("ec_low error")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	ecLow := float32(ecLowFloat)

	ecHighFloat, err := strconv.ParseFloat(r.PostFormValue("ec_high"), 32)
	if ecHighFloat < 0.1 || err != nil {
		fmt.Println("ec_high error")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	ecHigh := float32(ecHighFloat)

	plant := struct {
		Name                  string  `json:"name,omitempty"`
		CommonName            string  `json:"common_name"`
		SeedCompany           string  `json:"seed_company"`
		ExpectedDaysToHarvest int     `json:"expected_days_to_harvest"`
		Type                  string  `json:"type"`
		PhLow                 float32 `json:"ph_low,omitempty"`
		PhHigh                float32 `json:"ph_high,omitempty"`
		ECLow                 float32 `json:"ec_low,omitempty"`
		ECHigh                float32 `json:"ec_high,omitempty"`
	}{
		Name:                  name,
		CommonName:            common_name,
		SeedCompany:           seed_company,
		ExpectedDaysToHarvest: expected_days_to_harvest,
		Type:                  harvestType,
		PhLow:                 phLow,
		PhHigh:                phHigh,
		ECLow:                 ecLow,
		ECHigh:                ecHigh,
	}

	data, err := json.Marshal(plant)
	if err != nil {
		fmt.Println("Marshall error")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	req, _ := http.NewRequest("POST", app.plantList.Endpoint, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.do error")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Printf("unexpected status: %s", resp.Status)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/plant/create", http.StatusSeeOther)
}
