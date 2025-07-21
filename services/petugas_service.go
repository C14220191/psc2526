package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type PetugasService struct {
	DB *gorm.DB
}

func NewPetugasService(db *gorm.DB) *PetugasService {
	return &PetugasService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
