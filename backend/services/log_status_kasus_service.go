package services

import (
	"database/sql"
	"backend/models"
)

type LogStatusKasusService struct {
	DB *sql.DB
}

func NewLogStatusKasusService(db *sql.DB) *LogStatusKasusService {
	return &LogStatusKasusService{DB: db}
}

func (s *LogStatusKasusService) Create(data *models.LogStatusKasus) error {
	query := `
		INSERT INTO log_status_kasus 
		(kasus_id, deskripsi, waktu, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)`
	_, err := s.DB.Exec(query,
		data.KasusID, data.Deskripsi, data.Waktu, data.CreatedAt, data.UpdatedAt)
	return err
}

func (s *LogStatusKasusService) GetByID(id uint) (*models.LogStatusKasus, error) {
	query := `
		SELECT id, kasus_id, deskripsi, waktu, created_at, updated_at 
		FROM log_status_kasus WHERE id = ?`
	row := s.DB.QueryRow(query, id)

	var result models.LogStatusKasus
	err := row.Scan(&result.ID, &result.KasusID, &result.Deskripsi, &result.Waktu, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *LogStatusKasusService) Update(data *models.LogStatusKasus) error {
	query := `
		UPDATE log_status_kasus 
		SET deskripsi = ?, waktu = ?, updated_at = ? 
		WHERE id = ?`
	_, err := s.DB.Exec(query, data.Deskripsi, data.Waktu, data.UpdatedAt, data.ID)
	return err
}

func (s *LogStatusKasusService) Delete(id uint) error {
	query := `DELETE FROM log_status_kasus WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
