package services

import (
	"database/sql"
	"backend/models"
)

type PasienService struct {
	DB *sql.DB
}

func NewPasienService(db *sql.DB) *PasienService {
	return &PasienService{DB: db}
}

func (s *PasienService) Create(data *models.Pasien) error {
	return nil
}

func (s *PasienService) GetByID(id uint) (*models.Pasien, error) {
	return nil, nil
}

func (s *PasienService) Update(data *models.Pasien) error {
	return nil
}

func (s *PasienService) Delete(id uint) error {
	return nil
}
