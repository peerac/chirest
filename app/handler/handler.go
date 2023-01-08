package handler

import (
	"peerac/go-chi-rest-example/internal/data"

	"gorm.io/gorm"
)

type Handler struct {
	models data.Models
}

func NewHandler(db *gorm.DB) Handler {
	return Handler{
		models: data.NewModels(db),
	}
}
