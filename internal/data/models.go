package data

import (
	"gorm.io/gorm"
)

type Models struct {
	Movies MovieModel
}

func NewModels(db *gorm.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}
