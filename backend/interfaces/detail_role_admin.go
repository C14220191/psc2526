package interfaces
import "backend/models"

type DetailRoleAdminService interface {
	Create(data *models.DetailRoleAdmin) error
	GetByID(id uint) (*models.DetailRoleAdmin, error)
	Update(data *models.DetailRoleAdmin) error
	Delete(id uint) error
}

