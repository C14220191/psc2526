package services

import (
	"database/sql"
	"backend/models"
)

type KendaraanService struct {
	DB *sql.DB
}

func NewKendaraanService(db *sql.DB) *KendaraanService {
	return &KendaraanService{DB: db}
}

func (s *KendaraanService) Create(data *models.Kendaraan) error {
	// TODO: implement insert query
	return nil
}

func (s *KendaraanService) GetByID(id uint) (*models.Kendaraan, error) {
	// TODO: implement select by ID query
	return nil, nil
}

func (s *KendaraanService) Update(data *models.Kendaraan) error {
	// TODO: implement update query
	return nil
}

func (s *KendaraanService) Delete(id uint) error {
	// TODO: implement delete query
	return nil
}
