package intefaces
import "backend/models"

type LogAktivitasService interface {
	Create(data *models.LogAktivitas) error
	GetByID(id uint) (*models.LogAktivitas, error)
	Update(data *models.LogAktivitas) error
	Delete(id uint) error
}