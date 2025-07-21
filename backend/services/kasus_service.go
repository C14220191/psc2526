package services

import (
	"database/sql"
	"backend/models"
)

type KasusService struct {
	DB *sql.DB
}

func NewKasusService(db *sql.DB) *KasusService {
	return &KasusService{DB: db}
}

func (s *KasusService) Create(data *models.Kasus) error {
	query := `INSERT INTO kasus (
		judul, deskripsi, kategori_kasus_id,
		koordinat_latitude, koordinat_longitude, alamat_lengkap,
		status, waktu_kejadian, pelapor_id,
		created_at, updated_at
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := s.DB.Exec(query,
		data.Judul, data.Deskripsi, data.KategoriKasusID,
		data.KoordinatLatitude, data.KoordinatLongitude, data.AlamatLengkap,
		data.Status, data.WaktuKejadian, data.PelaporID,
		data.CreatedAt, data.UpdatedAt,
	)
	return err
}

func (s *KasusService) GetByID(id uint) (*models.Kasus, error) {
	query := `SELECT
		id, judul, deskripsi, kategori_kasus_id,
		koordinat_latitude, koordinat_longitude, alamat_lengkap,
		status, waktu_kejadian, pelapor_id,
		created_at, updated_at, deleted_at
		FROM kasus WHERE id = ?`

	row := s.DB.QueryRow(query, id)

	var result models.Kasus
	err := row.Scan(
		&result.ID, &result.Judul, &result.Deskripsi, &result.KategoriKasusID,
		&result.KoordinatLatitude, &result.KoordinatLongitude, &result.AlamatLengkap,
		&result.Status, &result.WaktuKejadian, &result.PelaporID,
		&result.CreatedAt, &result.UpdatedAt, &result.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *KasusService) Update(data *models.Kasus) error {
	query := `UPDATE kasus SET
		judul = ?, deskripsi = ?, kategori_kasus_id = ?,
		koordinat_latitude = ?, koordinat_longitude = ?, alamat_lengkap = ?,
		status = ?, waktu_kejadian = ?, pelapor_id = ?, updated_at = ?
		WHERE id = ?`

	_, err := s.DB.Exec(query,
		data.Judul, data.Deskripsi, data.KategoriKasusID,
		data.KoordinatLatitude, data.KoordinatLongitude, data.AlamatLengkap,
		data.Status, data.WaktuKejadian, data.PelaporID, data.UpdatedAt,
		data.ID,
	)
	return err
}

func (s *KasusService) Delete(id uint) error {
	query := `DELETE FROM kasus WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
