package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type LogStatusPetugasService struct {
	DB *gorm.DB
}

func NewLogStatusPetugasService(db *gorm.DB) *LogStatusPetugasService {
	return &LogStatusPetugasService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
