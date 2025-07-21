package services

import (
	"database/sql"
	"backend/models"
)

type LogTindakanPetugasService struct {
	DB *sql.DB
}

func NewLogTindakanPetugasService(db *sql.DB) *LogTindakanPetugasService {
	return &LogTindakanPetugasService{DB: db}
}

func (s *LogTindakanPetugasService) Create(data *models.LogTindakanPetugas) error {
	query := `
		INSERT INTO log_tindakan_petugas 
		(kasus_id, petugas_id, rincian, waktu_tindakan, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?)`
	_, err := s.DB.Exec(query,
		data.KasusID, data.PetugasID, data.Rincian,
		data.WaktuTindakan, data.CreatedAt, data.UpdatedAt,
	)
	return err
}

func (s *LogTindakanPetugasService) GetByID(id uint) (*models.LogTindakanPetugas, error) {
	query := `
		SELECT id, kasus_id, petugas_id, rincian, waktu_tindakan, created_at, updated_at
		FROM log_tindakan_petugas WHERE id = ?`
	row := s.DB.QueryRow(query, id)

	var result models.LogTindakanPetugas
	err := row.Scan(
		&result.ID, &result.KasusID, &result.PetugasID, &result.Rincian,
		&result.WaktuTindakan, &result.CreatedAt, &result.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *LogTindakanPetugasService) Update(data *models.LogTindakanPetugas) error {
	query := `
		UPDATE log_tindakan_petugas 
		SET rincian = ?, waktu_tindakan = ?, updated_at = ? 
		WHERE id = ?`
	_, err := s.DB.Exec(query, data.Rincian, data.WaktuTindakan, data.UpdatedAt, data.ID)
	return err
}

func (s *LogTindakanPetugasService) Delete(id uint) error {
	query := `DELETE FROM log_tindakan_petugas WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
