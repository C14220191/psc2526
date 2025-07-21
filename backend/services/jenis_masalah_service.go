package services

import (
	"database/sql"
	"backend/models"
)

type JenisMasalahService struct {
	DB *sql.DB
}

func NewJenisMasalahService(db *sql.DB) *JenisMasalahService {
	return &JenisMasalahService{DB: db}
}

func (s *JenisMasalahService) Create(data *models.JenisMasalah) error {
	return nil
}

func (s *JenisMasalahService) GetByID(id uint) (*models.JenisMasalah, error) {
	return nil, nil
}

func (s *JenisMasalahService) Update(data *models.JenisMasalah) error {
	return nil
}

func (s *JenisMasalahService) Delete(id uint) error {
	return nil
}
