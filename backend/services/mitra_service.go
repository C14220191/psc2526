package services

import (
	"database/sql"
	"backend/models"
)

type MitraService struct {
	DB *sql.DB
}

func NewMitraService(db *sql.DB) *MitraService {
	return &MitraService{DB: db}
}

func (s *MitraService) Create(data *models.Mitra) error {
	query := `INSERT INTO mitra (nama, alamat, kontak, email, created_at, updated_at)
	          VALUES (?, ?, ?, ?, ?, ?)`
	_, err := s.DB.Exec(query, data.Nama, data.Alamat, data.Kontak, data.Email, data.CreatedAt, data.UpdatedAt)
	return err
}

func (s *MitraService) GetByID(id uint) (*models.Mitra, error) {
	query := `SELECT id, nama, alamat, kontak, email, created_at, updated_at FROM mitra WHERE id = ?`
	row := s.DB.QueryRow(query, id)
	var result models.Mitra
	err := row.Scan(&result.ID, &result.Nama, &result.Alamat, &result.Kontak, &result.Email, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *MitraService) Update(data *models.Mitra) error {
	query := `UPDATE mitra SET nama = ?, alamat = ?, kontak = ?, email = ?, updated_at = ? WHERE id = ?`
	_, err := s.DB.Exec(query, data.Nama, data.Alamat, data.Kontak, data.Email, data.UpdatedAt, data.ID)
	return err
}

func (s *MitraService) Delete(id uint) error {
	query := `DELETE FROM mitra WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
