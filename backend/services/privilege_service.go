package services

import (
	"database/sql"
	"backend/models"
)

type PrivilegeService struct {
	DB *sql.DB
}

func NewPrivilegeService(db *sql.DB) *PrivilegeService {
	return &PrivilegeService{DB: db}
}

func (s *PrivilegeService) Create(data *models.Privilege) error {
	return nil
}

func (s *PrivilegeService) GetByID(id uint) (*models.Privilege, error) {
	return nil, nil
}

func (s *PrivilegeService) Update(data *models.Privilege) error {
	return nil
}

func (s *PrivilegeService) Delete(id uint) error {
	return nil
}
