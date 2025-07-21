package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type BeritaService struct {
	DB *gorm.DB
}

func NewBeritaService(db *gorm.DB) *BeritaService {
	return &BeritaService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
