package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type KategoriPelaporanService struct {
	DB *gorm.DB
}

func NewKategoriPelaporanService(db *gorm.DB) *KategoriPelaporanService {
	return &KategoriPelaporanService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
