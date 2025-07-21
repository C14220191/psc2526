package services

import (
	"database/sql"
	"backend/models"
)

type BeritaService struct {
	DB *sql.DB
}

func NewBeritaService(db *sql.DB) *BeritaService {
	return &BeritaService{DB: db}
}

func (s *BeritaService) Create(data *models.Berita) error {
	return nil
}

func (s *BeritaService) GetByID(id uint) (*models.Berita, error) {
	return nil, nil
}

func (s *BeritaService) Update(data *models.Berita) error {
	return nil
}

func (s *BeritaService) Delete(id uint) error {
	return nil
}
