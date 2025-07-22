package interfaces
import "backend/models"

type TindakanService interface {
    Create(data *models.Tindakan) error
    GetByID(id uint) (*models.Tindakan, error)
    Update(data *models.Tindakan) error
    Delete(id uint) error
}