package intefaces
import "backend/models"

type RoleService interface {
    Create(data *models.Role) error
    GetByID(id uint) (*models.Role, error)
    Update(data *models.Role) error
    Delete(id uint) error
}