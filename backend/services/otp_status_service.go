package services

import (
	"database/sql"
	"backend/models"
)

type OTPStatusService struct {
	DB *sql.DB
}

func NewOTPStatusService(db *sql.DB) *OTPStatusService {
	return &OTPStatusService{DB: db}
}

func (s *OTPStatusService) Create(data *models.OTPStatus) error {
	return nil
}

func (s *OTPStatusService) GetByID(id uint) (*models.OTPStatus, error) {
	return nil, nil
}

func (s *OTPStatusService) Update(data *models.OTPStatus) error {
	return nil
}

func (s *OTPStatusService) Delete(id uint) error {
	return nil
}
