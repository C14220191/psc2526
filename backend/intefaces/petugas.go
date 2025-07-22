package intefaces
import "backend/models"

type PetugasService interface {
    Create(data *models.Petugas) error
    GetByID(id uint) (*models.Petugas, error)
    Update(data *models.Petugas) error
    Delete(id uint) error
}
