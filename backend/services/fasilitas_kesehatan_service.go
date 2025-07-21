package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type FasilitasKesehatanService struct {
	DB *gorm.DB
}

func NewFasilitasKesehatanService(db *gorm.DB) *FasilitasKesehatanService {
	return &FasilitasKesehatanService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
