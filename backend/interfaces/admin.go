package interfaces

import (
	"backend/models"
	"context"
)

type AdminInterface interface {
	Create(ctx context.Context, data *models.AdminCreate) (*models.Response, error)
	GetAll(ctx context.Context, filter *models.AdminFilter) (*models.Response, error)
	GetByID(id uint) (*models.Admin, error)
	Update(data *models.Admin) error
	Delete(id uint) error
}