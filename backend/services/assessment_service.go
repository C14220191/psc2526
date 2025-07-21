package services

import (
	"database/sql"
	"backend/models"
)

type AssessmentService struct {
	DB *sql.DB
}

func NewAssessmentService(db *sql.DB) *AssessmentService {
	return &AssessmentService{DB: db}
}

func (s *AssessmentService) Create(data *models.Assessment) error {
	return nil
}

func (s *AssessmentService) GetByID(id uint) (*models.Assessment, error) {
	return nil, nil
}

func (s *AssessmentService) Update(data *models.Assessment) error {
	return nil
}

func (s *AssessmentService) Delete(id uint) error {
	return nil
}
