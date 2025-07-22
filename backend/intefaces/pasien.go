package intefaces
import "backend/models"

type PasienService interface {
	Create(data *models.Pasien) error
	GetByID(id uint) (*models.Pasien, error)
	Update(data *models.Pasien) error
	Delete(id uint) error
}