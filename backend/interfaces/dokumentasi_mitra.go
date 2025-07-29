package interfaces

import (
	"context"
	"backend/models"
)

type DokumentasiMitraInterface interface {
	Create(ctx context.Context, data *models.DokumentasiMitra) (*models.Response, error)
	GetByID(ctx context.Context, dokumentasi *models.DokumentasiMitra, id uint) (*models.Response, error)
	Update(ctx context.Context, data *models.DokumentasiMitra) (*models.Response, error)
	Delete(id uint) error
}
