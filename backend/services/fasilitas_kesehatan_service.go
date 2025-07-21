package services

import (
	"database/sql"
	"backend/models"
)

type FasilitasKesehatanService struct {
	DB *sql.DB
}

func NewFasilitasKesehatanService(db *sql.DB) *FasilitasKesehatanService {
	return &FasilitasKesehatanService{DB: db}
}

func (s *FasilitasKesehatanService) Create(data *models.FasilitasKesehatan) error {
	query := "INSERT INTO fasilitas_kesehatan (nama, alamat, tipe) VALUES (?, ?, ?)"
	_, err := s.DB.Exec(query, data.Nama, data.Alamat, data.Tipe)
	return err
}

func (s *FasilitasKesehatanService) GetByID(id uint) (*models.FasilitasKesehatan, error) {
	query := "SELECT id, nama, alamat, tipe FROM fasilitas_kesehatan WHERE id = ?"
	row := s.DB.QueryRow(query, id)

	var result models.FasilitasKesehatan
	err := row.Scan(&result.ID, &result.Nama, &result.Alamat, &result.Tipe)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *FasilitasKesehatanService) Update(data *models.FasilitasKesehatan) error {
	query := "UPDATE fasilitas_kesehatan SET nama = ?, alamat = ?, tipe = ? WHERE id = ?"
	_, err := s.DB.Exec(query, data.Nama, data.Alamat, data.Tipe, data.ID)
	return err
}

func (s *FasilitasKesehatanService) Delete(id uint) error {
	query := "DELETE FROM fasilitas_kesehatan WHERE id = ?"
	_, err := s.DB.Exec(query, id)
	return err
}
