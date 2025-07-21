package services

import (
	"database/sql"
	"backend/models"
)

type LogTindakanPetugasService struct {
	DB *sql.DB
}

func NewLogTindakanPetugasService(db *sql.DB) *LogTindakanPetugasService {
	return &LogTindakanPetugasService{DB: db}
}

func (s *LogTindakanPetugasService) Create(data *models.LogTindakanPetugas) error {
	return nil
}

func (s *LogTindakanPetugasService) GetByID(id uint) (*models.LogTindakanPetugas, error) {
	return nil, nil
}

func (s *LogTindakanPetugasService) Update(data *models.LogTindakanPetugas) error {
	return nil
}

func (s *LogTindakanPetugasService) Delete(id uint) error {
	return nil
}
