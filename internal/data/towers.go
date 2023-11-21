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
