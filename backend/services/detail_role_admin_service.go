package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type DetailRoleAdminService struct {
	DB *gorm.DB
}

func NewDetailRoleAdminService(db *gorm.DB) *DetailRoleAdminService {
	return &DetailRoleAdminService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
