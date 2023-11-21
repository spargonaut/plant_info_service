package data

import "database/sql"

type PlantProfileModels struct {
	Plants PlantProfile
}

type GrowTowerProfileModels struct {
	GrowTowers GrowTowerProfile
}

func NewPlantProfileModels(db *sql.DB) PlantProfileModels {
	return PlantProfileModels{
		Plants: PlantProfile{DB: db},
	}
}

func NewGrowTowerProfileModels(db *sql.DB) GrowTowerProfileModels {
	return GrowTowerProfileModels{
		GrowTowers: GrowTowerProfile{DB: db},
	}
}
