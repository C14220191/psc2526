package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type PrivilegeService struct {
	DB *gorm.DB
}

func NewPrivilegeService(db *gorm.DB) *PrivilegeService {
	return &PrivilegeService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
