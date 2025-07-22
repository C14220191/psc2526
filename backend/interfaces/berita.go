package interfaces
import "backend/models"

type BeritaService interface {
	Create(data *models.Berita) error
	GetByID(id uint) (*models.Berita, error)
	Update(data *models.Berita) error
	Delete(id uint) error
}