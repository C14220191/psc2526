package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type OtpStatusService struct {
	DB *gorm.DB
}

func NewOtpStatusService(db *gorm.DB) *OtpStatusService {
	return &OtpStatusService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
