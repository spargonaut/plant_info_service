package data

import (
	"database/sql"
	"fmt"
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

type PlantProfile struct {
	DB *sql.DB
}

func (p PlantProfile) Insert(plant *Plant) error {
	fmt.Println("attempting to insert the plant:")
	fmt.Println(plant)

	query := `
		INSERT INTO plants (name, common_name, seed_company, expected_days_to_harvest, type, ph_low, ph_high, ec_low, ec_high)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at, version`

	args := []interface{}{
		plant.Name,
		plant.CommonName,
		plant.SeedCompany,
		plant.ExpectedDaysToHarvest,
		plant.Type,
		plant.PhLow,
		plant.PhHigh,
		plant.ECLow,
		plant.ECHigh,
	}

	return p.DB.QueryRow(query, args...).Scan(&plant.ID, &plant.CreatedAt, &plant.Version)
}
