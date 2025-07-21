package services

import (
	"database/sql"
	"backend/models"
)

type DetailPrivilegeRoleService struct {
	DB *sql.DB
}

func NewDetailPrivilegeRoleService(db *sql.DB) *DetailPrivilegeRoleService {
	return &DetailPrivilegeRoleService{DB: db}
}

func (s *DetailPrivilegeRoleService) Create(data *models.DetailPrivilegeRole) error {
	query := "INSERT INTO detail_privilege_role (role_id, privilege_id, created_at, updated_at) VALUES (?, ?, ?, ?)"
	_, err := s.DB.Exec(query, data.RoleID, data.PrivilegeID, data.CreatedAt, data.UpdatedAt)
	return err
}

func (s *DetailPrivilegeRoleService) GetByID(id uint) (*models.DetailPrivilegeRole, error) {
	query := "SELECT id, role_id, privilege_id, created_at, updated_at FROM detail_privilege_role WHERE id = ?"
	row := s.DB.QueryRow(query, id)

	var result models.DetailPrivilegeRole
	err := row.Scan(&result.ID, &result.RoleID, &result.PrivilegeID, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *DetailPrivilegeRoleService) Update(data *models.DetailPrivilegeRole) error {
	query := "UPDATE detail_privilege_role SET role_id = ?, privilege_id = ?, updated_at = ? WHERE id = ?"
	_, err := s.DB.Exec(query, data.RoleID, data.PrivilegeID, data.UpdatedAt, data.ID)
	return err
}

func (s *DetailPrivilegeRoleService) Delete(id uint) error {
	query := "DELETE FROM detail_privilege_role WHERE id = ?"
	_, err := s.DB.Exec(query, id)
	return err
}
