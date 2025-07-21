package services

import (
	"database/sql"
	"backend/models"
)

type AdminService struct {
	DB *sql.DB
}

func NewAdminService(db *sql.DB) *AdminService {
	return &AdminService{DB: db}
}

func (s *AdminService) Create(data *models.Admin) error {
	query := `
		INSERT INTO admin (username, password, nama_lengkap, email, no_telepon, id_role, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := s.DB.Exec(query,
		data.Username, data.Password, data.NamaLengkap,
		data.Email, data.NoTelepon, data.IDRole,
		data.CreatedAt, data.UpdatedAt,
	)
	return err
}

func (s *AdminService) GetByID(id uint) (*models.Admin, error) {
	query := `
		SELECT id, username, password, nama_lengkap, email, no_telepon, id_role, created_at, updated_at
		FROM admin WHERE id = ?`
	row := s.DB.QueryRow(query, id)

	var result models.Admin
	err := row.Scan(
		&result.ID, &result.Username, &result.Password,
		&result.NamaLengkap, &result.Email, &result.NoTelepon,
		&result.IDRole, &result.CreatedAt, &result.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *AdminService) Update(data *models.Admin) error {
	query := `
		UPDATE admin SET username = ?, password = ?, nama_lengkap = ?, email = ?, no_telepon = ?, id_role = ?, updated_at = ?
		WHERE id = ?`
	_, err := s.DB.Exec(query,
		data.Username, data.Password, data.NamaLengkap,
		data.Email, data.NoTelepon, data.IDRole,
		data.UpdatedAt, data.ID,
	)
	return err
}

func (s *AdminService) Delete(id uint) error {
	query := `DELETE FROM admin WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
