package interfaces
import "backend/models"

type MitraService interface {
	Create(data *models.Mitra) error
	GetByID(id uint) (*models.Mitra, error)
	Update(data *models.Mitra) error
	Delete(id uint) error
}
