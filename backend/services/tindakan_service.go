package services

import (
	"database/sql"
	"backend/models"
)

type TindakanService struct {
	DB *sql.DB
}

func NewTindakanService(db *sql.DB) *TindakanService {
	return &TindakanService{DB: db}
}

func (s *TindakanService) Create(data *models.Tindakan) error {
	return nil
}

func (s *TindakanService) GetByID(id uint) (*models.Tindakan, error) {
	return nil, nil
}

func (s *TindakanService) Update(data *models.Tindakan) error {
	return nil
}

func (s *TindakanService) Delete(id uint) error {
	return nil
}
