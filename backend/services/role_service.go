package services

import (
	"database/sql"
	"backend/models"
)

type RoleService struct {
	DB *sql.DB
}

func NewRoleService(db *sql.DB) *RoleService {
	return &RoleService{DB: db}
}

func (s *RoleService) Create(data *models.Role) error {
	return nil
}

func (s *RoleService) GetByID(id uint) (*models.Role, error) {
	return nil, nil
}

func (s *RoleService) Update(data *models.Role) error {
	return nil
}

func (s *RoleService) Delete(id uint) error {
	return nil
}
