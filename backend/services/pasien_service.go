package services

import (
	"database/sql"
	"backend/models"
)

type PasienService struct {
	DB *sql.DB
}

func NewPasienService(db *sql.DB) *PasienService {
	return &PasienService{DB: db}
}

func (s *PasienService) Create(data *models.Pasien) error {
	query := `INSERT INTO pasien (nama, umur, jenis_kelamin, alamat,tempat lahir, tanggal lahir ,  created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := s.DB.Exec(query, data.Nama, data.Umur, data.JenisKelamin, data.Alamat, data.TempatLahir,data.TanggalLahir, data.CreatedAt, data.UpdatedAt)
	return err
}

func (s *PasienService) GetByID(id uint) (*models.Pasien, error) {
	query := `SELECT id, nama, umur, jenis_kelamin, alamat, tempat lahir, tanggal lahir, created_at, updated_at FROM pasien WHERE id = ?`
	row := s.DB.QueryRow(query, id)

	var result models.Pasien
	err := row.Scan(&result.ID, &result.Nama, &result.Umur, &result.JenisKelamin, &result.Alamat, &result.TempatLahir, result.TanggalLahir ,&result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *PasienService) Update(data *models.Pasien) error {
	query := `UPDATE pasien SET nama = ?, umur = ?, jenis_kelamin = ?, alamat = ?, tempat_lahir = ? , tanggal_lahir = ?, updated_at = ? WHERE id = ?`
	_, err := s.DB.Exec(query, data.Nama, data.Umur, data.JenisKelamin, data.Alamat, data.TempatLahir , data.TanggalLahir  ,data.UpdatedAt, data.ID)
	return err
}

func (s *PasienService) Delete(id uint) error {
	query := `DELETE FROM pasien WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
