package data

import "database/sql"

type PlantProfileModels struct {
	Plants PlantProfile
}

func NewPlantProfileModels(db *sql.DB) PlantProfileModels {
	return PlantProfileModels{
		Plants: PlantProfile{DB: db},
	}
}
