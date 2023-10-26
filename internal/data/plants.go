package data

import (
	"database/sql"
	"errors"
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

func (p PlantProfile) GetAll() ([]*Plant, error) {
	query := `
		SELECT *
		FROM plants
		ORDER BY id`

	rows, err := p.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	plants := []*Plant{}

	for rows.Next() {
		var plant Plant
		err := rows.Scan(
			&plant.ID,
			&plant.CreatedAt,
			&plant.Name,
			&plant.CommonName,
			&plant.SeedCompany,
			&plant.ExpectedDaysToHarvest,
			&plant.Type,
			&plant.PhLow,
			&plant.PhHigh,
			&plant.ECLow,
			&plant.ECHigh,
			&plant.Version,
		)
		if err != nil {
			return nil, err
		}

		plants = append(plants, &plant)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return plants, nil
}

func (p PlantProfile) Delete(id int64) error {
	if id < 1 {
		return errors.New("record not found")
	}

	query := `
		DELETE FROM plants
		WHERE id = $1`

	results, err := p.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}
