package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type PertanyaanService struct {
	DB *gorm.DB
}

func NewPertanyaanService(db *gorm.DB) *PertanyaanService {
	return &PertanyaanService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
