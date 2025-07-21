package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type KategoriKasusService struct {
	DB *gorm.DB
}

func NewKategoriKasusService(db *gorm.DB) *KategoriKasusService {
	return &KategoriKasusService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
