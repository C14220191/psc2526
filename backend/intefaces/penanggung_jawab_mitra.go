package intefaces
import "backend/models"

type PenanggungJawabMitraService interface {
	Create(data *models.PenanggungJawabMitra) error
	GetByID(id uint) (*models.PenanggungJawabMitra, error)
	Update(data *models.PenanggungJawabMitra) error
	Delete(id uint) error
}
