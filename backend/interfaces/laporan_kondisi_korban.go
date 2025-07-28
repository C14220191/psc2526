package interfaces

import (
	"context"
	"backend/models"
)

type LaporanKondisiKorbanInterface interface {
	Create(ctx context.Context, data *models.LaporanKondisiKorban) (*models.Response, error)
	GetByID(ctx context.Context, laporan *models.LaporanKondisiKorban, id uint) (*models.Response, error)
	Update(ctx context.Context, data *models.LaporanKondisiKorban) (*models.Response, error)
	Delete(id uint) error
}
