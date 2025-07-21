package services

import (
	"database/sql"
	"backend/models"
)

type PelaporService struct {
	DB *sql.DB
}

func NewPelaporService(db *sql.DB) *PelaporService {
	return &PelaporService{DB: db}
}

func (s *PelaporService) Create(data *models.Pelapor) error {
	query := `
		INSERT INTO pelapor (id_pelapor, nama_pelapor, lokasi_pelapor, no_telp, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)`
	_, err := s.DB.Exec(query,
		data.IDPelapor, data.NamaPelapor, data.LokasiPelapor, data.NoTelp,
		data.CreatedAt, data.UpdatedAt,
	)
	return err
}

func (s *PelaporService) GetByID(id int64) (*models.Pelapor, error) {
	query := `
		SELECT id, id_pelapor, nama_pelapor, lokasi_pelapor, no_telp, created_at, updated_at
		FROM pelapor WHERE id = ?`
	row := s.DB.QueryRow(query, id)

	var result models.Pelapor
	err := row.Scan(
		&result.ID, &result.IDPelapor, &result.NamaPelapor,
		&result.LokasiPelapor, &result.NoTelp,
		&result.CreatedAt, &result.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *PelaporService) Update(data *models.Pelapor) error {
	query := `
		UPDATE pelapor
		SET nama_pelapor = ?, lokasi_pelapor = ?, no_telp = ?, updated_at = ?
		WHERE id = ?`
	_, err := s.DB.Exec(query,
		data.NamaPelapor, data.LokasiPelapor, data.NoTelp,
		data.UpdatedAt, data.ID,
	)
	return err
}

func (s *PelaporService) Delete(id int64) error {
	query := `DELETE FROM pelapor WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
