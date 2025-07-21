package services

import (
	"database/sql"
	"backend/models"
)

type DetailRoleAdminService struct {
	DB *sql.DB
}

func NewDetailRoleAdminService(db *sql.DB) *DetailRoleAdminService {
	return &DetailRoleAdminService{DB: db}
}

func (s *DetailRoleAdminService) Create(data *models.DetailRoleAdmin) error {
	return nil
}

func (s *DetailRoleAdminService) GetByID(id uint) (*models.DetailRoleAdmin, error) {
	return nil, nil
}

func (s *DetailRoleAdminService) Update(data *models.DetailRoleAdmin) error {
	return nil
}

func (s *DetailRoleAdminService) Delete(id uint) error {
	return nil
}
