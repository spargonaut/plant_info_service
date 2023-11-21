package data

import (
	"database/sql"
	"fmt"
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

func (p GrowTowerProfile) Insert(growTower *GrowTower) error {
	fmt.Println("attempting to insert the grow tower:")
	fmt.Println(growTower.Name)

	query := `
		INSERT INTO growtowers (name, type, target_ph_low, target_ph_high, target_ec_low, target_ec_high)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, version`

	args := []interface{}{
		growTower.Name,
		growTower.Type,
		growTower.TargetPhLow,
		growTower.TargetPhHigh,
		growTower.TargetECLow,
		growTower.TargetECHigh,
	}

	return p.DB.QueryRow(query, args...).Scan(&growTower.ID, &growTower.CreatedAt, &growTower.Version)
}

func (gt GrowTowerProfile) GetAll() ([]*GrowTower, error) {
	query := `
		SELECT *
		FROM growtowers
		ORDER BY id`

	rows, err := gt.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	towers := []*GrowTower{}

	for rows.Next() {
		var tower GrowTower
		err := rows.Scan(
			&tower.ID,
			&tower.CreatedAt,
			&tower.Name,
			&tower.Type,
			&tower.TargetPhLow,
			&tower.TargetPhHigh,
			&tower.TargetECLow,
			&tower.TargetECHigh,
			&tower.Version,
		)
		if err != nil {
			return nil, err
		}

		towers = append(towers, &tower)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return towers, nil
}
