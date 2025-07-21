package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type JenisMasalahService struct {
	DB *gorm.DB
}

func NewJenisMasalahService(db *gorm.DB) *JenisMasalahService {
	return &JenisMasalahService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
