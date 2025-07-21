package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type LaporanObatService struct {
	DB *gorm.DB
}

func NewLaporanObatService(db *gorm.DB) *LaporanObatService {
	return &LaporanObatService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
