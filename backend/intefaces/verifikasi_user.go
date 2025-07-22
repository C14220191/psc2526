package intefaces

import "backend/models"

type VerifikasiUserService interface {
	Create(data *models.VerifikasiUser) error
	GetByID(id uint) (*models.VerifikasiUser, error)
	Update(data *models.VerifikasiUser) error
	Delete(id uint) error
}
