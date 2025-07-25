package services

import (
	"database/sql"
	"backend/models"
)

type PenanggungJawabMitraService struct {
	DB *sql.DB
}

func NewPenanggungJawabMitraService(db *sql.DB) *PenanggungJawabMitraService {
	return &PenanggungJawabMitraService{DB: db}
}

func (s *PenanggungJawabMitraService) Create(data *models.PenanggungJawabMitra) error {
	query := `INSERT INTO penanggung_jawab_mitra (id_mitra, nama, jabatan, kontak, email, created_at, updated_at)
	          VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := s.DB.Exec(query, data.IDMitra, data.Nama, data.Jabatan, data.Kontak, data.Email, data.CreatedAt, data.UpdatedAt)
	return err
}

func (s *PenanggungJawabMitraService) GetByID(id uint) (*models.PenanggungJawabMitra, error) {
	query := `SELECT id, id_mitra, nama, jabatan, kontak, email, created_at, updated_at FROM penanggung_jawab_mitra WHERE id = ?`
	row := s.DB.QueryRow(query, id)
	var result models.PenanggungJawabMitra
	err := row.Scan(&result.ID, &result.IDMitra, &result.Nama, &result.Jabatan, &result.Kontak, &result.Email, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *PenanggungJawabMitraService) Update(data *models.PenanggungJawabMitra) error {
	query := `UPDATE penanggung_jawab_mitra SET nama = ?, jabatan = ?, kontak = ?, email = ?, updated_at = ? WHERE id = ?`
	_, err := s.DB.Exec(query, data.Nama, data.Jabatan, data.Kontak, data.Email, data.UpdatedAt, data.ID)
	return err
}

func (s *PenanggungJawabMitraService) Delete(id uint) error {
	query := `DELETE FROM penanggung_jawab_mitra WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
