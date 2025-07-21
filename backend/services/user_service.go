package services

import (
	"gorm.io/gorm"
	"your_project/models"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

// Implement service methods like Create, GetByID, Update, Delete
