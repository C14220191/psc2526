package services

import (
	"database/sql"
	"backend/models"
)

type KategoriKasusService struct {
	DB *sql.DB
}

func NewKategoriKasusService(db *sql.DB) *KategoriKasusService {
	return &KategoriKasusService{DB: db}
}

func (s *KategoriKasusService) Create(data *models.KategoriKasus) error {
	return nil
}

func (s *KategoriKasusService) GetByID(id uint) (*models.KategoriKasus, error) {
	return nil, nil
}

func (s *KategoriKasusService) Update(data *models.KategoriKasus) error {
	return nil
}

func (s *KategoriKasusService) Delete(id uint) error {
	return nil
}
