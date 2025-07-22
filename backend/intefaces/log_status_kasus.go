package intefaces
import "backend/models"

type LogStatusKasusService interface {
	Create(data *models.LogStatusKasus) error
	GetByID(id uint) (*models.LogStatusKasus, error)
	Update(data *models.LogStatusKasus) error
	Delete(id uint) error
}
