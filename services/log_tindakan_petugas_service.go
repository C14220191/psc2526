package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type LogTindakanPetugasService struct {
	DB *gorm.DB
}

func NewLogTindakanPetugasService(db *gorm.DB) *LogTindakanPetugasService {
	return &LogTindakanPetugasService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
