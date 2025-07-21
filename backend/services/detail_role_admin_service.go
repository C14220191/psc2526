package services

import (
	"database/sql"
	"backend/models"
)

type DetailRoleAdminService struct {
	DB *sql.DB
}

func NewDetailRoleAdminService(db *sql.DB) *DetailRoleAdminService {
	return &DetailRoleAdminService{DB: db}
}

func (s *DetailRoleAdminService) Create(data *models.DetailRoleAdmin) error {
	query := "INSERT INTO detail_role_admin (id_admin, id_role, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?, ?)"
	_, err := s.DB.Exec(query, data.IDAdmin, data.IDRole, data.CreatedAt, data.UpdatedAt, data.DeletedAt)
	return err
}

func (s *DetailRoleAdminService) GetByID(id uint) (*models.DetailRoleAdmin, error) {
	query := "SELECT id_detail_role, id_admin, id_role, created_at, updated_at, deleted_at FROM detail_role_admin WHERE id_detail_role = ?"
	row := s.DB.QueryRow(query, id)

	var result models.DetailRoleAdmin
	err := row.Scan(&result.ID, &result.IDAdmin, &result.IDRole, &result.CreatedAt, &result.UpdatedAt, &result.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *DetailRoleAdminService) Update(data *models.DetailRoleAdmin) error {
	query := "UPDATE detail_role_admin SET id_admin = ?, id_role = ?, updated_at = ?, deleted_at = ? WHERE id_detail_role = ?"
	_, err := s.DB.Exec(query, data.IDAdmin, data.IDRole, data.UpdatedAt, data.DeletedAt, data.ID)
	return err
}

func (s *DetailRoleAdminService) Delete(id uint) error {
	query := "DELETE FROM detail_role_admin WHERE id_detail_role = ?"
	_, err := s.DB.Exec(query, id)
	return err
}
