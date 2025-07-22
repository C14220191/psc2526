package intefaces
import "backend/models"

type KategoriKasusService interface {
	Create(data *models.KategoriKasus) error
	GetByID(id uint) (*models.KategoriKasus, error)
	Update(data *models.KategoriKasus) error
	Delete(id uint) error
}