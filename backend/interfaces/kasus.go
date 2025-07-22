package interfaces
import "backend/models"

type KasusService interface {
	Create(data *models.Kasus) error
	GetByID(id uint) (*models.Kasus, error)
	Update(data *models.Kasus) error
	Delete(id uint) error
}