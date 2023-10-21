package data

import (
	"time"
)

type Plant struct {
	ID                    int64     `json:"id"`
	CreatedAt             time.Time `json:"-"`
	Name                  string    `json:"name,omitempty"`
	CommonName            string    `json:"common_name"`
	SeedCompany           string    `json:"seed_company"`
	ExpectedDaysToHarvest int32     `json:"expected_days_to_harvest"`
	Type                  string    `json:"type"`
	PhLow                 float32   `json:"ph_low,omitempty"`
	PhHigh                float32   `json:"ph_high,omitempty"`
	ECLow                 float32   `json:"ec_low,omitempty"`
	ECHigh                float32   `json:"ec_high,omitempty"`
	Version               int32     `json:"-"`
}
