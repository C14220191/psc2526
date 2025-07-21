package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type AssessmentService struct {
	DB *gorm.DB
}

func NewAssessmentService(db *gorm.DB) *AssessmentService {
	return &AssessmentService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
