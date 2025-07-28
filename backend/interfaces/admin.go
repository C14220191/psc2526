package interfaces

import (
	"backend/models"
	"context"
)

type AdminInterface interface {
	Create(ctx context.Context, data *models.AdminCreate) (*models.Response, error)
	GetAll(ctx context.Context, filter *models.AdminFilter) (*models.Response, error)
	GetByID(ctx context.Context,admin *models.Admin, id uint) (*models.Response, error)
	Update(ctx context.Context, data *models.AdminUpdate) (*models.Response, error)
	Delete(id uint) error
}