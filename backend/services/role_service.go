package services

import (
	"database/sql"
	"backend/models"
)

type RoleService struct {
	DB *sql.DB
}

func NewRoleService(db *sql.DB) *RoleService {
	return &RoleService{DB: db}
}

func (s *RoleService) Create(data *models.Role) error {
	query := `INSERT INTO role (namaRole, created_at, updated_at, deleted_at) VALUES (?, ?, ?)`
	_, err := s.DB.Exec(query, data.NamaRole, data.CreatedAt, data.UpdatedAt , data.DeletedAt)
	return err
}

func (s *RoleService) GetByID(id uint) (*models.Role, error) {
	query := `SELECT id, namaRole, created_at, updated_at FROM role WHERE id = ?`
	row := s.DB.QueryRow(query, id)

	var result models.Role
	err := row.Scan(&result.ID, &result.NamaRole, &result.CreatedAt, &result.UpdatedAt, result.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *RoleService) Update(data *models.Role) error {
	query := `UPDATE role SET namaRole = ?, updated_at = ? WHERE id = ?`
	_, err := s.DB.Exec(query, data.NamaRole, data.UpdatedAt, data.ID, data.DeletedAt)
	return err
}

func (s *RoleService) Delete(id uint) error {
	query := `DELETE FROM role WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
