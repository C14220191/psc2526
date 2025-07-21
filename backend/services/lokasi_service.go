package services

import (
	"database/sql"
	"backend/models"
)

type LokasiService struct {
	DB *sql.DB
}

func NewLokasiService(db *sql.DB) *LokasiService {
	return &LokasiService{DB: db}
}

func (s *LokasiService) Create(data *models.Lokasi) error {
	query := `INSERT INTO lokasi 
		(nama_lokasi, alamat, latitude, longitude, tipe, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := s.DB.Exec(query,
		data.NamaLokasi, data.Alamat, data.Latitude, data.Longitude, data.Tipe, data.CreatedAt, data.UpdatedAt)
	return err
}

func (s *LokasiService) GetByID(id uint) (*models.Lokasi, error) {
	query := `SELECT id, nama_lokasi, alamat, latitude, longitude, tipe, created_at, updated_at 
	          FROM lokasi WHERE id = ?`
	row := s.DB.QueryRow(query, id)

	var result models.Lokasi
	err := row.Scan(&result.ID, &result.NamaLokasi, &result.Alamat, &result.Latitude, &result.Longitude, &result.Tipe, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *LokasiService) Update(data *models.Lokasi) error {
	query := `UPDATE lokasi 
		SET nama_lokasi = ?, alamat = ?, latitude = ?, longitude = ?, tipe = ?, updated_at = ? 
		WHERE id = ?`
	_, err := s.DB.Exec(query,
		data.NamaLokasi, data.Alamat, data.Latitude, data.Longitude, data.Tipe, data.UpdatedAt, data.ID)
	return err
}

func (s *LokasiService) Delete(id uint) error {
	query := `DELETE FROM lokasi WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
