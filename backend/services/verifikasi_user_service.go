package services

import (
	"database/sql"
	"backend/models"
)

type VerifikasiUserService struct {
	DB *sql.DB
}

func NewVerifikasiUserService(db *sql.DB) *VerifikasiUserService {
	return &VerifikasiUserService{DB: db}
}

func (s *VerifikasiUserService) Create(data *models.VerifikasiUser) error {
	return nil
}

func (s *VerifikasiUserService) GetByID(id uint) (*models.VerifikasiUser, error) {
	return nil, nil
}

func (s *VerifikasiUserService) Update(data *models.VerifikasiUser) error {
	return nil
}

func (s *VerifikasiUserService) Delete(id uint) error {
	return nil
}
