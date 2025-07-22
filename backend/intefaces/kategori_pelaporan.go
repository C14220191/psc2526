package intefaces
import "backend/models"

type KategoriPelaporanService interface {
	Create(data *models.KategoriPelaporan) error
	GetByID(id uint) (*models.KategoriPelaporan, error)
	Update(data *models.KategoriPelaporan) error
	Delete(id uint) error
}