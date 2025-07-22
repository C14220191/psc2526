package interfaces
import "backend/models"

type DokterService interface {
	Create(data *models.Dokter) error
	GetByID(id uint) (*models.Dokter, error)
	Update(data *models.Dokter) error
	Delete(id uint) error
}
