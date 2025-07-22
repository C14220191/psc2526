package interfaces
import "backend/models"

type LokasiService interface {
	Create(data *models.Lokasi) error
	GetByID(id uint) (*models.Lokasi, error)
	Update(data *models.Lokasi) error
	Delete(id uint) error
}
