package services

import (
	"database/sql"
	"backend/models"
)

type KendaraanService struct {
	DB *sql.DB
}

func NewKendaraanService(db *sql.DB) *KendaraanService {
	return &KendaraanService{DB: db}
}

func (s *KendaraanService) Create(data *models.Kendaraan) error {
	query := `INSERT INTO kendaraan (nomor_plat, jenis, status, merek, warna, tahun, created_at, updated_at)
	          VALUES (?, ?, ?, ?, ?)`
	_, err := s.DB.Exec(query, data.NomorPlat, data.Jenis, data.Status,data.Merek,data.Warna,data.Tahun, data.CreatedAt, data.UpdatedAt)
	return err
}

func (s *KendaraanService) GetByID(id uint) (*models.Kendaraan, error) {
	query := `SELECT id, nomor_plat, jenis, status, created_at, updated_at FROM kendaraan WHERE id = ?`
	row := s.DB.QueryRow(query, id)
	var result models.Kendaraan
	err := row.Scan(&result.ID, &result.NomorPlat, &result.Jenis, &result.Status, &result.Merek,&result.Warna,&result.Tahun, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *KendaraanService) Update(data *models.Kendaraan) error {
	query := `UPDATE kendaraan SET nomor_plat = ?, jenis = ?, status = ?, updated_at = ? WHERE id = ?`
	_, err := s.DB.Exec(query, data.NomorPlat, data.Jenis, data.Status,data.Merek,data.Warna,data.Tahun, data.CreatedAt, data.UpdatedAt)
	return err
}

func (s *KendaraanService) Delete(id uint) error {
	query := `DELETE FROM kendaraan WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
