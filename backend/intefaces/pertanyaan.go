package intefaces
import "backend/models"

type PertanyaanService interface {
	Create(data *models.Pertanyaan) error
	GetByID(id uint) (*models.Pertanyaan, error)
	Update(data *models.Pertanyaan) error
	Delete(id uint) error
}
