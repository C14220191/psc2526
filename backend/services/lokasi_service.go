package services

import (
	"database/sql"
	"backend/models"
)

type LokasiService struct {
	DB *sql.DB
}

func NewLokasiService(db *sql.DB) *LokasiService {
	return &LokasiService{DB: db}
}

func (s *LokasiService) Create(data *models.Lokasi) error {
	return nil
}

func (s *LokasiService) GetByID(id uint) (*models.Lokasi, error) {
	return nil, nil
}

func (s *LokasiService) Update(data *models.Lokasi) error {
	return nil
}

func (s *LokasiService) Delete(id uint) error {
	return nil
}
