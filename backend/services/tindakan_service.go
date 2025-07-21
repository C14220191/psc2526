package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type TindakanService struct {
	DB *gorm.DB
}

func NewTindakanService(db *gorm.DB) *TindakanService {
	return &TindakanService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
