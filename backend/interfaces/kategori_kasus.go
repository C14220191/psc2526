package interfaces

import "backend/models"

type KategoriKasusInterface interface {
	Create(data *models.KategoriKasus) error
	GetAll() ([]*models.KategoriKasus, error)
	GetByID(id uint) (*models.KategoriKasus, error)
	Update(data *models.KategoriKasus) error
	Delete(id uint) error
}
