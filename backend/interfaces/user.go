package interfaces
import "backend/models"

type UserService interface {
	Create(data *models.User) error
	GetByID(id int) (*models.User, error)
	Update(data *models.User) error
	Delete(id int) error
}
