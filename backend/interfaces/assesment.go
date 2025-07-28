package interfaces

import (
	"backend/models"
	"context"
)

type AssessmentInterface interface {
	Create(data *models.AssessmentCreate, ctx context.Context) (*models.Response,error)
	GetAll(ctx context.Context, data models.AssessmentGetAllResponse)(*models.Response, error)
	GetByID(ctx context.Context, assessment models.Assessment, id uint) (*models.Response, error)
	Update(data *models.AssessmentUpdate, ctx context.Context) (*models.Response, error)
	Delete(id uint, ctx context.Context) (*models.Response, error)
}
