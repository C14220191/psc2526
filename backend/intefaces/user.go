package intefaces
import "backend/models"

type userService interface {
	Create(data *models.User) error
	GetByID(id uint) (*models.User, error)
	Update(data *models.User) error
	Delete(id uint) error
}
