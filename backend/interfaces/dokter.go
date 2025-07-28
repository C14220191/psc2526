package interfaces

import (
	"backend/models"
	"context"
)

type DokterInterface interface {
	Create(ctx context.Context, data *models.DokterCreate) (*models.Response, error)
	GetAll(ctx context.Context, filter *models.DokterFilter) (*models.Response, error)
	GetByID(ctx context.Context, dokter *models.Dokter, id uint) (*models.Response, error)
	Update(ctx context.Context, data *models.DokterUpdate) (*models.Response, error)
	Delete(id uint) error
}
