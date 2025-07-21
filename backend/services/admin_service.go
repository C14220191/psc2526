package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type AdminService struct {
	DB *gorm.DB
}

func NewAdminService(db *gorm.DB) *AdminService {
	return &AdminService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
