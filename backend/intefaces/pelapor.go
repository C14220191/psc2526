package intefaces
import "backend/models"

type PelaporService interface {
	Create(data *models.Pelapor) error
	GetByID(id uint) (*models.Pelapor, error)
	Update(data *models.Pelapor) error
	Delete(id uint) error
}