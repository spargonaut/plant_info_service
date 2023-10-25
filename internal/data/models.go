package data

import "database/sql"

type Models struct {
	Plants PlantProfile
}

func NewModels(db *sql.DB) Models {
	return Models{
		Plants: PlantProfile{DB: db},
	}
}
