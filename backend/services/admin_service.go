package services

import (
	"backend/interfaces"
	"backend/models"
	"context"
	"database/sql"
	"net/http"
)

type AdminService struct {
	DB *sql.DB
	_cfg any
}

var _ interfaces.AdminInterface = &AdminService{}
func NewAdminService(db *sql.DB) *AdminService {
	return &AdminService{DB: db}
}

func (s *AdminService) Create(ctx context.Context, data *models.AdminCreate) (*models.Response, error) {
	var res models.Response

	tx, err := s.DB.BeginTx(ctx, &sql.TxOptions{})

	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to begin transaction"
		res.Data = nil
		return nil, nil
	}

	query := `
		INSERT INTO admin (username, password, nama_lengkap, email, no_telepon, id_role)
		VALUES (?, ?, ?, ?, ?, ?)`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		tx.Rollback()
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to prepare statement"
		res.Data = nil
		return &res, nil
	}
	defer stmt.Close()


	result, err := s.DB.ExecContext(ctx, query,
		data.Username, data.Password, data.NamaLengkap,
		data.Email, data.NoTelepon, data.IDRole,)

	if err != nil {
		tx.Rollback()
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to execute query"
		res.Data = nil
		return &res, nil
	}

	lastID, _ := result.LastInsertId()

	if err := tx.Commit(); err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to commit transaction"
		res.Data = nil
		return &res, nil
	}

	res.StatusCode = http.StatusCreated
	res.Message = "Admin created successfully"
	res.Data = lastID
	return &res, nil
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
