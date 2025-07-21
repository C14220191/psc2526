package services

import (
	"database/sql"
	"backend/models"
)

type KategoriPelaporanService struct {
	DB *sql.DB
}

func NewKategoriPelaporanService(db *sql.DB) *KategoriPelaporanService {
	return &KategoriPelaporanService{DB: db}
}

func (s *KategoriPelaporanService) Create(data *models.KategoriPelaporan) error {
	query := `INSERT INTO kategori_pelaporan (nama, deskripsi, created_at, updated_at) VALUES (?, ?, ?, ?)`
	_, err := s.DB.Exec(query, data.Nama, data.Deskripsi, data.CreatedAt, data.UpdatedAt)
	return err
}

func (s *KategoriPelaporanService) GetByID(id uint) (*models.KategoriPelaporan, error) {
	query := `SELECT id, nama, deskripsi, created_at, updated_at FROM kategori_pelaporan WHERE id = ?`
	row := s.DB.QueryRow(query, id)

	var result models.KategoriPelaporan
	err := row.Scan(&result.ID, &result.Nama, &result.Deskripsi, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *KategoriPelaporanService) Update(data *models.KategoriPelaporan) error {
	query := `UPDATE kategori_pelaporan SET nama = ?, deskripsi = ?, updated_at = ? WHERE id = ?`
	_, err := s.DB.Exec(query, data.Nama, data.Deskripsi, data.UpdatedAt, data.ID)
	return err
}

func (s *KategoriPelaporanService) Delete(id uint) error {
	query := `DELETE FROM kategori_pelaporan WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
