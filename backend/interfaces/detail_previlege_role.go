package interfaces
import "backend/models"

type DetailPrivilegeRoleService interface {
	Create(data *models.DetailPrivilegeRole) error
	GetByID(id uint) (*models.DetailPrivilegeRole, error)
	Update(data *models.DetailPrivilegeRole) error
	Delete(id uint) error
}

