package services

import (
	"database/sql"
	"backend/models"
)

type LaporanObatService struct {
	DB *sql.DB
}

func NewLaporanObatService(db *sql.DB) *LaporanObatService {
	return &LaporanObatService{DB: db}
}

func (s *LaporanObatService) Create(data *models.LaporanObat) error {
	return nil
}

func (s *LaporanObatService) GetByID(id uint) (*models.LaporanObat, error) {
	return nil, nil
}

func (s *LaporanObatService) Update(data *models.LaporanObat) error {
	return nil
}

func (s *LaporanObatService) Delete(id uint) error {
	return nil
}
