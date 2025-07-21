package services

import (
	"database/sql"
	"backend/models"
)

type KategoriKasusService struct {
	DB *sql.DB
}

func NewKategoriKasusService(db *sql.DB) *KategoriKasusService {
	return &KategoriKasusService{DB: db}
}

func (s *KategoriKasusService) Create(data *models.KategoriKasus) error {
	query := `INSERT INTO kategori_kasus (nama_kategori, created_at, updated_at) VALUES (?, ?, ?)`
	_, err := s.DB.Exec(query, data.Nama_kategori, data.CreatedAt, data.UpdatedAt)
	return err
}

func (s *KategoriKasusService) GetByID(id uint) (*models.KategoriKasus, error) {
	query := `SELECT id, nama_kategori, created_at, updated_at FROM kategori_kasus WHERE id = ?`
	row := s.DB.QueryRow(query, id)

	var result models.KategoriKasus
	err := row.Scan(&result.ID, &result.Nama_kategori, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *KategoriKasusService) Update(data *models.KategoriKasus) error {
	query := `UPDATE kategori_kasus SET nama_kategori = ?, updated_at = ? WHERE id = ?`
	_, err := s.DB.Exec(query, data.Nama_kategori, data.UpdatedAt, data.ID)
	return err
}

func (s *KategoriKasusService) Delete(id uint) error {
	query := `DELETE FROM kategori_kasus WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
