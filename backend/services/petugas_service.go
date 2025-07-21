package services

import (
	"database/sql"
	"backend/models"
)

type PetugasService struct {
	DB *sql.DB
}

func NewPetugasService(db *sql.DB) *PetugasService {
	return &PetugasService{DB: db}
}

func (s *PetugasService) Create(data *models.Petugas) error {
	return nil
}

func (s *PetugasService) GetByID(id uint) (*models.Petugas, error) {
	return nil, nil
}

func (s *PetugasService) Update(data *models.Petugas) error {
	return nil
}

func (s *PetugasService) Delete(id uint) error {
	return nil
}
