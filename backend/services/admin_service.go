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
	query := "INSERT INTO admin (username, password, nama, email, id_role) VALUES (?, ?, ?, ?, ?)"
	_, err := s.DB.Exec(query, data.Username, data.Password, data.Nama, data.Email, data.IDRole)
	return err
}

func (s *AdminService) GetByID(id uint) (*models.Admin, error) {
	query := "SELECT id, username, password, nama, email, id_role FROM admin WHERE id = ?"
	row := s.DB.QueryRow(query, id)

	var result models.Admin
	err := row.Scan(&result.ID, &result.Username, &result.Password, &result.Nama, &result.Email, &result.IDRole)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *AdminService) Update(data *models.Admin) error {
	query := "UPDATE admin SET username = ?, password = ?, nama = ?, email = ?, id_role = ? WHERE id = ?"
	_, err := s.DB.Exec(query, data.Username, data.Password, data.Nama, data.Email, data.IDRole, data.ID)
	return err
}

func (s *AdminService) Delete(id uint) error {
	query := "DELETE FROM admin WHERE id = ?"
	_, err := s.DB.Exec(query, id)
	return err
}
