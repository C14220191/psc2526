package services

import (
	"database/sql"
	"backend/models"
)

type VerifikasiUserService struct {
	DB *sql.DB
}

func NewVerifikasiUserService(db *sql.DB) *VerifikasiUserService {
	return &VerifikasiUserService{DB: db}
}

func (s *VerifikasiUserService) Create(data *models.VerifikasiUser) error {
	query := `INSERT INTO verifikasi_user (waktu_pengajuan, status, created_at, updated_at)
	          VALUES (?, ?, ?, ?)`
	_, err := s.DB.Exec(query, data.WaktuPengajuan, data.Status, data.CreatedAt, data.UpdatedAt)
	return err
}

func (s *VerifikasiUserService) GetByID(id uint) (*models.VerifikasiUser, error) {
	query := `SELECT id, waktu_pengajuan, status, created_at, updated_at FROM verifikasi_user WHERE id = ?`
	row := s.DB.QueryRow(query, id)

	var result models.VerifikasiUser
	err := row.Scan(&result.ID, &result.WaktuPengajuan, &result.Status, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *VerifikasiUserService) Update(data *models.VerifikasiUser) error {
	query := `UPDATE verifikasi_user SET status = ?, updated_at = ? WHERE id = ?`
	_, err := s.DB.Exec(query, data.Status, data.UpdatedAt, data.ID)
	return err
}

func (s *VerifikasiUserService) Delete(id uint) error {
	query := `DELETE FROM verifikasi_user WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
