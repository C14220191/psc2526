package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type LokasiService struct {
	DB *gorm.DB
}

func NewLokasiService(db *gorm.DB) *LokasiService {
	return &LokasiService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
