package interfaces
import "backend/models"

type KendaraanService interface {
	Create(data *models.Kendaraan) error
	GetByID(id uint) (*models.Kendaraan, error)
	Update(data *models.Kendaraan) error
	Delete(id uint) error
}
