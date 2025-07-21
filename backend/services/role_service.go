package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type RoleService struct {
	DB *gorm.DB
}

func NewRoleService(db *gorm.DB) *RoleService {
	return &RoleService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
