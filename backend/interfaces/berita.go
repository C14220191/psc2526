package interfaces

import (
	"backend/models"
	"context"
)

type BeritaInterface interface {
	Create(data *models.BeritaCreate, ctx context.Context) (*models.Response, error)
	GetByID(id uint, ctx context.Context) (*models.Response, error)
	Update(data *models.BeritaUpdate, ctx context.Context, id uint) (*models.Response, error)
	Delete(id uint, ctx context.Context) (*models.Response, error)
}