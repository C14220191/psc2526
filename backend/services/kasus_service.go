package services

import (
	"database/sql"
	"backend/models"
)

type KasusService struct {
	DB *sql.DB
}

func NewKasusService(db *sql.DB) *KasusService {
	return &KasusService{DB: db}
}

func (s *KasusService) Create(data *models.Kasus) error {
	return nil
}

func (s *KasusService) GetByID(id uint) (*models.Kasus, error) {
	return nil, nil
}

func (s *KasusService) Update(data *models.Kasus) error {
	return nil
}

func (s *KasusService) Delete(id uint) error {
	return nil
}
