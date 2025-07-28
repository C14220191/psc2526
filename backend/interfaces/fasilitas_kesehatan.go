package interfaces

import (
	"context"
	"backend/models"
)

type FasilitasKesehatanInterface interface {
	Create(ctx context.Context, data *models.FasilitasKesehatanCreate) (*models.Response, error)
	GetAll(ctx context.Context, filter *models.FasilitasKesehatanFilter) (*models.Response, error)
	GetByID(ctx context.Context, fasilitas *models.FasilitasKesehatan, id uint) (*models.Response, error)
	Update(ctx context.Context, data *models.FasilitasKesehatanUpdate) (*models.Response, error)
	Delete(id uint) error
}
