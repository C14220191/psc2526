package services

import (
	"database/sql"
	"backend/models"
)

type DetailPrivilegeRoleService struct {
	DB *sql.DB
}

func NewDetailPrivilegeRoleService(db *sql.DB) *DetailPrivilegeRoleService {
	return &DetailPrivilegeRoleService{DB: db}
}

func (s *DetailPrivilegeRoleService) Create(data *models.DetailPrivilegeRole) error {
	return nil
}

func (s *DetailPrivilegeRoleService) GetByID(id uint) (*models.DetailPrivilegeRole, error) {
	return nil, nil
}

func (s *DetailPrivilegeRoleService) Update(data *models.DetailPrivilegeRole) error {
	return nil
}

func (s *DetailPrivilegeRoleService) Delete(id uint) error {
	return nil
}
