package services

import (
	"database/sql"
	"backend/models"
)

type JenisMasalahService struct {
	DB *sql.DB
}

func NewJenisMasalahService(db *sql.DB) *JenisMasalahService {
	return &JenisMasalahService{DB: db}
}

func (s *JenisMasalahService) Create(data *models.JenisMasalah) error {
	query := `INSERT INTO jenis_masalah (nama, deskripsi, created_at, updated_at) VALUES (?, ?, ?, ?)`
	_, err := s.DB.Exec(query, data.Nama, data.Deskripsi, data.CreatedAt, data.UpdatedAt)
	return err
}

func (s *JenisMasalahService) GetByID(id uint) (*models.JenisMasalah, error) {
	query := `SELECT id, nama, deskripsi, created_at, updated_at FROM jenis_masalah WHERE id = ?`
	row := s.DB.QueryRow(query, id)

	var result models.JenisMasalah
	err := row.Scan(&result.ID, &result.Nama, &result.Deskripsi, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *JenisMasalahService) Update(data *models.JenisMasalah) error {
	query := `UPDATE jenis_masalah SET nama = ?, deskripsi = ?, updated_at = ? WHERE id = ?`
	_, err := s.DB.Exec(query, data.Nama, data.Deskripsi, data.UpdatedAt, data.ID)
	return err
}

func (s *JenisMasalahService) Delete(id uint) error {
	query := `DELETE FROM jenis_masalah WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
