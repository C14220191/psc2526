package services

import (
	"database/sql"
	"backend/models"
)

type LogStatusPetugasService struct {
	DB *sql.DB
}

func NewLogStatusPetugasService(db *sql.DB) *LogStatusPetugasService {
	return &LogStatusPetugasService{DB: db}
}

func (s *LogStatusPetugasService) Create(data *models.LogStatusPetugas) error {
	query := `INSERT INTO log_status_petugas (petugas_id, status, waktu, timestamp, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := s.DB.Exec(query, data.PetugasID, data.Status, data.Waktu, data.Timestamp, data.CreatedAt, data.UpdatedAt)
	return err
}

func (s *LogStatusPetugasService) GetByID(id uint) (*models.LogStatusPetugas, error) {
	query := `SELECT id, petugas_id, status, waktu, timestamp, created_at, updated_at FROM log_status_petugas WHERE id = ?`
	row := s.DB.QueryRow(query, id)
	var result models.LogStatusPetugas
	err := row.Scan(&result.ID, &result.PetugasID, &result.Status, &result.Waktu, &result.Timestamp, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *LogStatusPetugasService) Update(data *models.LogStatusPetugas) error {
	query := `UPDATE log_status_petugas SET status = ?, waktu = ?, timestamp = ?, updated_at = ? WHERE id = ?`
	_, err := s.DB.Exec(query, data.Status, data.Waktu, data.Timestamp, data.UpdatedAt, data.ID)
	return err
}

func (s *LogStatusPetugasService) Delete(id uint) error {
	query := `DELETE FROM log_status_petugas WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
