package services

import (
	"database/sql"
	"backend/models"
)

type LogStatusKasusService struct {
	DB *sql.DB
}

func NewLogStatusKasusService(db *sql.DB) *LogStatusKasusService {
	return &LogStatusKasusService{DB: db}
}

func (s *LogStatusKasusService) Create(data *models.LogStatusKasus) error {
	return nil
}

func (s *LogStatusKasusService) GetByID(id uint) (*models.LogStatusKasus, error) {
	return nil, nil
}

func (s *LogStatusKasusService) Update(data *models.LogStatusKasus) error {
	return nil
}

func (s *LogStatusKasusService) Delete(id uint) error {
	return nil
}
