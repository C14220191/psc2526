package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type KendaraanService struct {
	DB *gorm.DB
}

func NewKendaraanService(db *gorm.DB) *KendaraanService {
	return &KendaraanService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
