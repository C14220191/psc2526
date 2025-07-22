package interfaces

import "backend/models"

type LaporanKondisiKorbanService interface {
	Create(data *models.LaporanKondisiKorban) error
	GetByID(id uint) (*models.LaporanKondisiKorban, error)
	Update(data *models.LaporanKondisiKorban) error
	Delete(id uint) error
}
