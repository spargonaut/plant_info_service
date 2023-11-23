package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Println("http method error")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/favicon.ico" && r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	plants, err := app.plantInfo.GetAll()
	if err != nil {
		fmt.Println("error getting all plants")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", plants)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (app *application) towerHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Println("http method error")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/favicon.ico" && r.URL.Path != "/towers" {
		http.NotFound(w, r)
		return
	}

	towers, err := app.towerInfo.GetAll()
	if err != nil {
		fmt.Println("error getting all towers")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/towerlist.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", towers)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
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
	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/create.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
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

	req, _ := http.NewRequest("POST", app.plantInfo.CommandEndpoint, bytes.NewBuffer(data))
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

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) deletePlant(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.deletePlantForm(w, r)
	case http.MethodPost:
		app.deletePlantProcess(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (app *application) deletePlantForm(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/delete.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (app *application) deletePlantProcess(w http.ResponseWriter, r *http.Request) {
	fmt.Println("processing the delete plant form")
	plantId := r.PostFormValue("id")
	if plantId == "" {
		fmt.Println("plantId error")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	req, _ := http.NewRequest("DELETE", app.plantInfo.CommandEndpoint+"/"+plantId, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.do error")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("unexpected status: %s when attempting to delete ID: %s", resp.Status, plantId)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
