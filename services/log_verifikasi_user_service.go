package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type LogVerifikasiUserService struct {
	DB *gorm.DB
}

func NewLogVerifikasiUserService(db *gorm.DB) *LogVerifikasiUserService {
	return &LogVerifikasiUserService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
