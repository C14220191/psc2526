package interfaces
import "backend/models"

type FasilitasKesehatanService interface {
	Create(data *models.FasilitasKesehatan) error
	GetByID(id uint) (*models.FasilitasKesehatan, error)
	Update(data *models.FasilitasKesehatan) error
	Delete(id uint) error
}