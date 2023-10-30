package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Plant struct {
	ID                    int64   `json:"id"`
	Name                  string  `json:"name,omitempty"`
	CommonName            string  `json:"common_name"`
	SeedCompany           string  `json:"seed_company"`
	ExpectedDaysToHarvest int32   `json:"expected_days_to_harvest"`
	Type                  string  `json:"type"`
	PhLow                 float32 `json:"ph_low,omitempty"`
	PhHigh                float32 `json:"ph_high,omitempty"`
	ECLow                 float32 `json:"ec_low,omitempty"`
	ECHigh                float32 `json:"ec_high,omitempty"`
}

type PlantsResponse struct {
	Plants *[]Plant `json:"plants"`
}

type PlantProfile struct {
	CommandEndpoint string
	QueryEndpoint   string
}

func (m *PlantProfile) GetAll() (*[]Plant, error) {
	resp, err := http.Get(m.QueryEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var plantsResp PlantsResponse
	err = json.Unmarshal(data, &plantsResp)
	if err != nil {
		return nil, err
	}

	return plantsResp.Plants, nil
}
