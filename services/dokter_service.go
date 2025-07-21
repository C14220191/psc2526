package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type DokterService struct {
	DB *gorm.DB
}

func NewDokterService(db *gorm.DB) *DokterService {
	return &DokterService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
