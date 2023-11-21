package data

import (
	"database/sql"
	"time"
)

type GrowTower struct {
	ID           int64     `json:"id"`
	CreatedAt    time.Time `json:"-"`
	Name         string    `json:"name,omitempty"`
	Type         string    `json:"type"`
	TargetPhLow  float32   `json:"target_ph_low,omitempty"`
	TargetPhHigh float32   `json:"target_ph_high,omitempty"`
	TargetECLow  float32   `json:"target_ec_low,omitempty"`
	TargetECHigh float32   `json:"target_ec_high,omitempty"`
	Version      int32     `json:"-"`
}

type GrowTowerProfile struct {
	DB *sql.DB
}
