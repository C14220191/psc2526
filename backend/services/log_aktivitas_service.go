package services

import (
	"database/sql"
	"backend/models"
)

type LogAktivitasService struct {
	DB *sql.DB
}

func NewLogAktivitasService(db *sql.DB) *LogAktivitasService {
	return &LogAktivitasService{DB: db}
}

func (s *LogAktivitasService) Create(data *models.LogAktivitas) error {
	query := `
	INSERT INTO log_aktivitas 
	(id_admin, id_petugas, aksi, deskripsi, waktu, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := s.DB.Exec(query,
		data.IDAdmin, data.IDPetugas, data.Aksi, data.Deskripsi,
		data.Waktu, data.CreatedAt, data.UpdatedAt)
	return err
}

func (s *LogAktivitasService) GetByID(id uint) (*models.LogAktivitas, error) {
	query := `
	SELECT id, id_admin, id_petugas, aksi, deskripsi, waktu, created_at, updated_at 
	FROM log_aktivitas WHERE id = ?`
	row := s.DB.QueryRow(query, id)

	var result models.LogAktivitas
	err := row.Scan(
		&result.ID, &result.IDAdmin, &result.IDPetugas, &result.Aksi,
		&result.Deskripsi, &result.Waktu, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *LogAktivitasService) Update(data *models.LogAktivitas) error {
	query := `
	UPDATE log_aktivitas 
	SET aksi = ?, deskripsi = ?, waktu = ?, updated_at = ?
	WHERE id = ?`
	_, err := s.DB.Exec(query,
		data.Aksi, data.Deskripsi, data.Waktu, data.UpdatedAt, data.ID)
	return err
}

func (s *LogAktivitasService) Delete(id uint) error {
	query := `DELETE FROM log_aktivitas WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
