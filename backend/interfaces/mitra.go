// interfaces/mitra_interface.go
package interfaces

import (
	"context"
	"backend/models"
)

type MitraService interface {
	Create(ctx context.Context, data *models.MitraCreate) (*models.Response, error)
	GetByID(ctx context.Context, id uint) (*models.Response, error)
	Update(ctx context.Context, data *models.MitraUpdate) (*models.Response, error)
	Delete(ctx context.Context, id uint) (*models.Response, error)
	GetAll(ctx context.Context, filter *models.MitraFilter) (*models.Response, error)
}
