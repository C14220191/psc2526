package services

import (
	"database/sql"
	"backend/models"
)

type LaporanKondisiKorbanService struct {
	DB *sql.DB
}

func NewLaporanKondisiKorbanService(db *sql.DB) *LaporanKondisiKorbanService {
	return &LaporanKondisiKorbanService{DB: db}
}

func (s *LaporanKondisiKorbanService) Create(data *models.LaporanKondisiKorban) error {
	return nil
}

func (s *LaporanKondisiKorbanService) GetByID(id uint) (*models.LaporanKondisiKorban, error) {
	return nil, nil
}

func (s *LaporanKondisiKorbanService) Update(data *models.LaporanKondisiKorban) error {
	return nil
}

func (s *LaporanKondisiKorbanService) Delete(id uint) error {
	return nil
}
