package services

import (
	"database/sql"
	"backend/models"
)

type UserService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) Create(data *models.User) error {
	return nil
}

func (s *UserService) GetByID(id uint) (*models.User, error) {
	return nil, nil
}

func (s *UserService) Update(data *models.User) error {
	return nil
}

func (s *UserService) Delete(id uint) error {
	return nil
}
