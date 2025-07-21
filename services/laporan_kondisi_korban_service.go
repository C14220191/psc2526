package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type LaporanKondisiKorbanService struct {
	DB *gorm.DB
}

func NewLaporanKondisiKorbanService(db *gorm.DB) *LaporanKondisiKorbanService {
	return &LaporanKondisiKorbanService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
