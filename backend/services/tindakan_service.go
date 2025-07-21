package services

import (
	"database/sql"
	"backend/models"
)

type TindakanService struct {
	DB *sql.DB
}

func NewTindakanService(db *sql.DB) *TindakanService {
	return &TindakanService{DB: db}
}

func (s *TindakanService) Create(data *models.Tindakan) error {
	query := `INSERT INTO tindakan (kasus_id, petugas_id, jenis, rincian, waktu, created_at, updated_at)
	          VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := s.DB.Exec(query, data.KasusID, data.PetugasID, data.Jenis, data.Rincian, data.Waktu, data.CreatedAt, data.UpdatedAt)
	return err
}

func (s *TindakanService) GetByID(id uint) (*models.Tindakan, error) {
	query := `SELECT id, kasus_id, petugas_id, jenis, rincian, waktu, created_at, updated_at 
	          FROM tindakan WHERE id = ?`
	row := s.DB.QueryRow(query, id)

	var result models.Tindakan
	err := row.Scan(&result.ID, &result.KasusID, &result.PetugasID, &result.Jenis, &result.Rincian, &result.Waktu, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *TindakanService) Update(data *models.Tindakan) error {
	query := `UPDATE tindakan 
	          SET jenis = ?, rincian = ?, waktu = ?, updated_at = ? 
	          WHERE id = ?`
	_, err := s.DB.Exec(query, data.Jenis, data.Rincian, data.Waktu, data.UpdatedAt, data.ID)
	return err
}

func (s *TindakanService) Delete(id uint) error {
	query := `DELETE FROM tindakan WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
