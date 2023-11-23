package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GrowTower struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name,omitempty"`
	Type         string  `json:"type"`
	TargetPhLow  float32 `json:"target_ph_low,omitempty"`
	TargetPhHigh float32 `json:"target_ph_high,omitempty"`
	TargetECLow  float32 `json:"target_ec_low,omitempty"`
	TargetECHigh float32 `json:"target_ec_high,omitempty"`
}

type GrowTowerResponse struct {
	GrowTowers *[]GrowTower `json:"grow_towers"`
}

type GrowTowerProfile struct {
	CommandEndpoint string
	QueryEndpoint   string
}

func (gtp *GrowTowerProfile) GetAll() (*[]GrowTower, error) {
	resp, err := http.Get(gtp.QueryEndpoint)
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

	var growTowersResp GrowTowerResponse
	err = json.Unmarshal(data, &growTowersResp)
	if err != nil {
		return nil, err
	}

	return growTowersResp.GrowTowers, nil
}
