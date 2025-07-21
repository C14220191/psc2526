package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type DetailPrivilegeRoleService struct {
	DB *gorm.DB
}

func NewDetailPrivilegeRoleService(db *gorm.DB) *DetailPrivilegeRoleService {
	return &DetailPrivilegeRoleService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
