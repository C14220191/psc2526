package interfaces

import (
	"backend/models"
)

type AdminService interface {
	Create(data *models.Admin) error
	GetByID(id uint) (*models.Admin, error)
	Update(data *models.Admin) error
	Delete(id uint) error
}