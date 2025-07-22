package interfaces
import "backend/models"

type JenisMasalahService interface {
	Create(data *models.JenisMasalah) error
	GetByID(id uint) (*models.JenisMasalah, error)
	Update(data *models.JenisMasalah) error
	Delete(id uint) error
}