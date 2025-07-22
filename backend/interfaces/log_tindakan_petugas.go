package interfaces
import "backend/models"

type LogTindakanPetugasService interface {
	Create(data *models.LogTindakanPetugas) error
	GetByID(id uint) (*models.LogTindakanPetugas, error)
	Update(data *models.LogTindakanPetugas) error
	Delete(id uint) error
}