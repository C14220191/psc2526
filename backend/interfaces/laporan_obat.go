package interfaces
import "backend/models"

type LaporanObatService interface {
	Create(data *models.LaporanObat) error
	GetByID(id uint) (*models.LaporanObat, error)
	Update(data *models.LaporanObat) error
	Delete(id uint) error
}
