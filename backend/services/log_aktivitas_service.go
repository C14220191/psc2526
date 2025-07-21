package services

import (
	"database/sql"
	"backend/models"
)

type LogAktivitasService struct {
	DB *sql.DB
}

func NewLogAktivitasService(db *sql.DB) *LogAktivitasService {
	return &LogAktivitasService{DB: db}
}

func (s *LogAktivitasService) Create(data *models.LogAktivitas) error {
	return nil
}

func (s *LogAktivitasService) GetByID(id uint) (*models.LogAktivitas, error) {
	return nil, nil
}

func (s *LogAktivitasService) Update(data *models.LogAktivitas) error {
	return nil
}

func (s *LogAktivitasService) Delete(id uint) error {
	return nil
}
