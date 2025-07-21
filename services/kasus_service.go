package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type KasusService struct {
	DB *gorm.DB
}

func NewKasusService(db *gorm.DB) *KasusService {
	return &KasusService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
