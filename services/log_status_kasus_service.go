package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type LogStatusKasusService struct {
	DB *gorm.DB
}

func NewLogStatusKasusService(db *gorm.DB) *LogStatusKasusService {
	return &LogStatusKasusService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
