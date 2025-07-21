package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type PelaporService struct {
	DB *gorm.DB
}

func NewPelaporService(db *gorm.DB) *PelaporService {
	return &PelaporService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
