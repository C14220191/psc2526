package interfaces
import "backend/models"

type LogStatusPetugasService interface {
	Create(data *models.LogStatusPetugas) error
	GetByID(id uint) (*models.LogStatusPetugas, error)
	Update(data *models.LogStatusPetugas) error
	Delete(id uint) error
}