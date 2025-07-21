package services

import (
	"database/sql"
	"backend/models"
)

type PertanyaanService struct {
	DB *sql.DB
}

func NewPertanyaanService(db *sql.DB) *PertanyaanService {
	return &PertanyaanService{DB: db}
}

func (s *PertanyaanService) Create(data *models.Pertanyaan) error {
	return nil
}

func (s *PertanyaanService) GetByID(id uint) (*models.Pertanyaan, error) {
	return nil, nil
}

func (s *PertanyaanService) Update(data *models.Pertanyaan) error {
	return nil
}

func (s *PertanyaanService) Delete(id uint) error {
	return nil
}
