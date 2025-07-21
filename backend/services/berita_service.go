package services

import (
	"database/sql"
	"backend/models"
)

type BeritaService struct {
	DB *sql.DB
}

func NewBeritaService(db *sql.DB) *BeritaService {
	return &BeritaService{DB: db}
}

func (s *BeritaService) Create(data *models.Berita) error {
	query := `
	INSERT INTO berita (id_admin_pembuat_berita, judul, isi, thumbnail, tanggal_date, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	_, err := s.DB.Exec(query,
		data.IDAdminPembuatBerita,
		data.Judul,
		data.Isi,
		data.Thumbnail,
		data.Tanggal,
		data.CreatedAt,
		data.UpdatedAt,
	)
	return err
}

func (s *BeritaService) GetByID(id uint) (*models.Berita, error) {
	query := `
	SELECT id_berita, id_admin_pembuat_berita, judul, isi, thumbnail, tanggal_date, created_at, updated_at, deleted_at
	FROM berita
	WHERE id_berita = ?
	`
	row := s.DB.QueryRow(query, id)

	var result models.Berita
	err := row.Scan(
		&result.ID,
		&result.IDAdminPembuatBerita,
		&result.Judul,
		&result.Isi,
		&result.Thumbnail,
		&result.Tanggal,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *BeritaService) Update(data *models.Berita) error {
	query := `
	UPDATE berita
	SET judul = ?, isi = ?, thumbnail = ?, tanggal_date = ?, updated_at = ?
	WHERE id_berita = ?
	`
	_, err := s.DB.Exec(query,
		data.Judul,
		data.Isi,
		data.Thumbnail,
		data.Tanggal,
		data.UpdatedAt,
		data.ID,
	)
	return err
}

func (s *BeritaService) Delete(id uint) error {
	query := `
	UPDATE berita
	SET deleted_at = ?
	WHERE id_berita = ?
	`
	_, err := s.DB.Exec(query, sql.NullTime{Valid: true, Time: sql.NullTime{}.Time}, id)
	return err
}
