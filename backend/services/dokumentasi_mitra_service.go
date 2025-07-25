package services

import (
	"database/sql"
	"backend/models"
)

type DokumentasiMitraService struct {
	DB *sql.DB
}

func NewDokumentasiMitraService(db *sql.DB) *DokumentasiMitraService {
	return &DokumentasiMitraService{DB: db}
}

func (s *DokumentasiMitraService) Create(data *models.DokumentasiMitra) error {
	query := `INSERT INTO dokumentasi_mitra (id_mitra, file_url, keterangan, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	_, err := s.DB.Exec(query, data.IDMitra, data.FileURL, data.Keterangan, data.CreatedAt, data.UpdatedAt)
	return err
}

func (s *DokumentasiMitraService) GetByID(id uint) (*models.DokumentasiMitra, error) {
	query := `SELECT id, id_mitra, file_url, keterangan, created_at, updated_at FROM dokumentasi_mitra WHERE id = ?`
	row := s.DB.QueryRow(query, id)

	var result models.DokumentasiMitra
	err := row.Scan(&result.ID, &result.IDMitra, &result.FileURL, &result.Keterangan, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *DokumentasiMitraService) Update(data *models.DokumentasiMitra) error {
	query := `UPDATE dokumentasi_mitra SET file_url = ?, keterangan = ?, updated_at = ? WHERE id = ?`
	_, err := s.DB.Exec(query, data.FileURL, data.Keterangan, data.UpdatedAt, data.ID)
	return err
}

func (s *DokumentasiMitraService) Delete(id uint) error {
	query := `DELETE FROM dokumentasi_mitra WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
