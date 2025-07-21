package services

import (
	"database/sql"
	"backend/models"
)

type PelaporService struct {
	DB *sql.DB
}

func NewPelaporService(db *sql.DB) *PelaporService {
	return &PelaporService{DB: db}
}

func (s *PelaporService) Create(data *models.Pelapor) error {
	return nil
}

func (s *PelaporService) GetByID(id uint) (*models.Pelapor, error) {
	return nil, nil
}

func (s *PelaporService) Update(data *models.Pelapor) error {
	return nil
}

func (s *PelaporService) Delete(id uint) error {
	return nil
}
