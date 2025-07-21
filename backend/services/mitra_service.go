package services

import (
	"database/sql"
	"backend/models"
)

type MitraService struct {
	DB *sql.DB
}

func NewMitraService(db *sql.DB) *MitraService {
	return &MitraService{DB: db}
}

func (s *MitraService) Create(data *models.Mitra) error {
	query := "INSERT INTO mitra (nama, alamat, ...) VALUES (?, ?, ...)"
	_, err := s.DB.Exec(query, data.Nama, data.Alamat /*, ...*/)
	return err
}

func (s *MitraService) GetByID(id uint) (*models.Mitra, error) {
	query := "SELECT id, nama, alamat FROM mitra WHERE id = ?"
	row := s.DB.QueryRow(query, id)

	var mitra models.Mitra
	err := row.Scan(&mitra.ID, &mitra.Nama, &mitra.Alamat /*, ...*/)
	if err != nil {
		return nil, err
	}
	return &mitra, nil
}

func (s *MitraService) Update(data *models.Mitra) error {
	query := "UPDATE mitra SET nama = ?, alamat = ? WHERE id = ?"
	_, err := s.DB.Exec(query, data.Nama, data.Alamat, data.ID)
	return err
}

func (s *MitraService) Delete(id uint) error {
	query := "DELETE FROM mitra WHERE id = ?"
	_, err := s.DB.Exec(query, id)
	return err
}
