package services

import (
	"database/sql"
	"backend/models"
)

type DokterService struct {
	DB *sql.DB
}

func NewDokterService(db *sql.DB) *DokterService {
	return &DokterService{DB: db}
}

func (s *DokterService) Create(data *models.Dokter) error {
	query := `
	INSERT INTO dokter (nama, spesialisasi, faskes_id, no_str, jenis_kelamin, kontak, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := s.DB.Exec(query, data.Nama, data.Spesialisasi, data.FaskesID, data.NoSTR, data.JenisKelamin, data.Kontak, data.CreatedAt, data.UpdatedAt)
	return err
}

func (s *DokterService) GetByID(id uint) (*models.Dokter, error) {
	query := `
	SELECT id, nama, spesialisasi, faskes_id, no_str, jenis_kelamin, kontak, created_at, updated_at
	FROM dokter WHERE id = ?`
	row := s.DB.QueryRow(query, id)

	var result models.Dokter
	err := row.Scan(&result.ID, &result.Nama, &result.Spesialisasi, &result.FaskesID, &result.NoSTR, &result.JenisKelamin, &result.Kontak, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *DokterService) Update(data *models.Dokter) error {
	query := `
	UPDATE dokter SET nama = ?, spesialisasi = ?, faskes_id = ?, no_str = ?, jenis_kelamin = ?, kontak = ?, updated_at = ?
	WHERE id = ?`
	_, err := s.DB.Exec(query, data.Nama, data.Spesialisasi, data.FaskesID, data.NoSTR, data.JenisKelamin, data.Kontak, data.UpdatedAt, data.ID)
	return err
}

func (s *DokterService) Delete(id uint) error {
	query := "DELETE FROM dokter WHERE id = ?"
	_, err := s.DB.Exec(query, id)
	return err
}
