package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type DokumentasiMitraService struct {
	DB *gorm.DB
}

func NewDokumentasiMitraService(db *gorm.DB) *DokumentasiMitraService {
	return &DokumentasiMitraService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
