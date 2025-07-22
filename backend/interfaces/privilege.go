package interfaces

import "backend/models"

type PrivilegeService interface {
	Create(data *models.Privilege) error
	GetByID(id uint) (*models.Privilege, error)
	Update(data *models.Privilege) error
	Delete(id uint) error
}
