package services

import (
	"database/sql"
	"backend/models"
)

type OtpStatusService struct {
	DB *sql.DB
}

func NewOtpStatusService(db *sql.DB) *OtpStatusService {
	return &OtpStatusService{DB: db}
}

func (s *OtpStatusService) Create(data *models.OTPStatus) error {
	query := `INSERT INTO otp_status (user_id, kode_otp, status, terkirim_at, expired_at, created_at, updated_at)
	          VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := s.DB.Exec(query,
		data.UserID, data.KodeOTP, data.Status,
		data.TerkirimAt, data.ExpiredAt,
		data.CreatedAt, data.UpdatedAt,
	)
	return err
}

func (s *OtpStatusService) GetByID(id uint) (*models.OTPStatus, error) {
	query := `SELECT id, user_id, kode_otp, status, terkirim_at, expired_at, created_at, updated_at FROM otp_status WHERE id = ?`
	row := s.DB.QueryRow(query, id)

	var result models.OTPStatus
	err := row.Scan(
		&result.ID, &result.UserID, &result.KodeOTP, &result.Status,
		&result.TerkirimAt, &result.ExpiredAt, &result.CreatedAt, &result.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *OtpStatusService) Update(data *models.OTPStatus) error {
	query := `UPDATE otp_status SET kode_otp = ?, status = ?, terkirim_at = ?, expired_at = ?, updated_at = ? WHERE id = ?`
	_, err := s.DB.Exec(query,
		data.KodeOTP, data.Status, data.TerkirimAt, data.ExpiredAt, data.UpdatedAt, data.ID,
	)
	return err
}

func (s *OtpStatusService) Delete(id uint) error {
	query := `DELETE FROM otp_status WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
