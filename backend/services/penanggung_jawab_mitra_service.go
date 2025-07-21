package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type PenanggungJawabMitraService struct {
	DB *gorm.DB
}

func NewPenanggungJawabMitraService(db *gorm.DB) *PenanggungJawabMitraService {
	return &PenanggungJawabMitraService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
