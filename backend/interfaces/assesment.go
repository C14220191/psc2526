package interfaces
import "backend/models"

type AssessmentService interface {
	Create(data *models.Assessment) error
	GetByID(id uint) (*models.Assessment, error)
	Update(data *models.Assessment) error
	Delete(id uint) error
}
