package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type PasienService struct {
	DB *gorm.DB
}

func NewPasienService(db *gorm.DB) *PasienService {
	return &PasienService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
