package intefaces
import "backend/models"

type DokumentasiMitraService interface {
	Create(data *models.DokumentasiMitra) error
	GetByID(id uint) (*models.DokumentasiMitra, error)
	Update(data *models.DokumentasiMitra) error
	Delete(id uint) error
}