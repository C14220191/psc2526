package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type MitraService struct {
	DB *gorm.DB
}

func NewMitraService(db *gorm.DB) *MitraService {
	return &MitraService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
