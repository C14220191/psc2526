package services

import (
	"database/sql"
	"backend/models"
)

type PetugasService struct {
	DB *sql.DB
}

func NewPetugasService(db *sql.DB) *PetugasService {
	return &PetugasService{DB: db}
}

func (s *PetugasService) Create(data *models.Petugas) error {
	query := `INSERT INTO petugas 
	(id_faskes, id_tindakan_petugas, nama, jenis_kelamin, profesi, username, password, no_hp, keahlian)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := s.DB.Exec(query,
		data.IDFaskes, data.IDTindakanPetugas, data.Nama, data.JenisKelamin, data.Profesi,
		data.Username, data.Password, data.NoHP, data.Keahlian)
	return err
}

func (s *PetugasService) GetByID(id uint) (*models.Petugas, error) {
	query := `SELECT id, id_faskes, id_tindakan_petugas, nama, jenis_kelamin, profesi, username, password, no_hp, keahlian 
	          FROM petugas WHERE id = ?`
	row := s.DB.QueryRow(query, id)

	var result models.Petugas
	err := row.Scan(&result.ID, &result.IDFaskes, &result.IDTindakanPetugas,
		&result.Nama, &result.JenisKelamin, &result.Profesi, &result.Username,
		&result.Password, &result.NoHP, &result.Keahlian)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *PetugasService) Update(data *models.Petugas) error {
	query := `UPDATE petugas 
	          SET id_faskes = ?, id_tindakan_petugas = ?, nama = ?, jenis_kelamin = ?, 
	              profesi = ?, username = ?, password = ?, no_hp = ?, keahlian = ?
	          WHERE id = ?`
	_, err := s.DB.Exec(query,
		data.IDFaskes, data.IDTindakanPetugas, data.Nama, data.JenisKelamin,
		data.Profesi, data.Username, data.Password, data.NoHP, data.Keahlian, data.ID)
	return err
}

func (s *PetugasService) Delete(id uint) error {
	query := `DELETE FROM petugas WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
