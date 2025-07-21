package services

import (
	"database/sql"
	"backend/models"
)

type KategoriPelaporanService struct {
	DB *sql.DB
}

func NewKategoriPelaporanService(db *sql.DB) *KategoriPelaporanService {
	return &KategoriPelaporanService{DB: db}
}

func (s *KategoriPelaporanService) Create(data *models.KategoriPelaporan) error {
	return nil
}

func (s *KategoriPelaporanService) GetByID(id uint) (*models.KategoriPelaporan, error) {
	return nil, nil
}

func (s *KategoriPelaporanService) Update(data *models.KategoriPelaporan) error {
	return nil
}

func (s *KategoriPelaporanService) Delete(id uint) error {
	return nil
}
