package services

import (
	"database/sql"
	"backend/models"
)

type PrivilegeService struct {
	DB *sql.DB
}

func NewPrivilegeService(db *sql.DB) *PrivilegeService {
	return &PrivilegeService{DB: db}
}

func (s *PrivilegeService) Create(data *models.Privilege) error {
	query := `INSERT INTO privilege (nama_privilege, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?)`
_, err := s.DB.Exec(query, data.NamaPrivilege, data.CreatedAt, data.UpdatedAt, data.DeletedAt)

	return err
}

func (s *PrivilegeService) GetByID(id uint) (*models.Privilege, error) {
	query := `SELECT id, nama_privilege, created_at, updated_at, deleted_at FROM privilege WHERE id = ?`
	row := s.DB.QueryRow(query, id)

	var result models.Privilege
	err := row.Scan(&result.ID, &result.NamaPrivilege, &result.CreatedAt, &result.UpdatedAt, &result.DeletedAt)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *PrivilegeService) Update(data *models.Privilege) error {
	query := `UPDATE privilege SET nama = ?, deskripsi = ?, updated_at = ? WHERE id = ?`
_, err := s.DB.Exec(query, data.NamaPrivilege, data.CreatedAt, data.UpdatedAt, data.DeletedAt)
	return err
}

func (s *PrivilegeService) Delete(id uint) error {
	query := `DELETE FROM privilege WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
