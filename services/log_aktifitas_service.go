package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type LogAktifitasService struct {
	DB *gorm.DB
}

func NewLogAktifitasService(db *gorm.DB) *LogAktifitasService {
	return &LogAktifitasService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
