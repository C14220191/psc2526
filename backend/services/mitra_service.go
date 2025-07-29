// services/mitra_service.go
package services

import (
	"backend/models"
	"context"
	"database/sql"
	"net/http"
	"strings"
	"time"
)

type MitraService struct {
	DB *sql.DB
}

func NewMitraService(db *sql.DB) *MitraService {
	return &MitraService{DB: db}
}

func (s *MitraService) Create(ctx context.Context, data *models.MitraCreate) (*models.Response, error) {
	var res models.Response
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to begin transaction"
		return &res, err
	}

	now := time.Now()
	query := `INSERT INTO mitra (nama, alamat, kontak, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`
	result, err := tx.ExecContext(ctx, query, data.Nama, data.Alamat, data.Kontak, data.Email, now, now)
	if err != nil {
		tx.Rollback()
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to insert mitra"
		return &res, err
	}

	id, _ := result.LastInsertId()
	tx.Commit()
	res.StatusCode = http.StatusCreated
	res.Message = "Mitra created successfully"
	res.Data = id
	return &res, nil
}

func (s *MitraService) GetByID(ctx context.Context, id uint) (*models.Response, error) {
	var res models.Response
	query := `SELECT id, nama, alamat, kontak, email, created_at, updated_at FROM mitra WHERE id = ?`
	row := s.DB.QueryRowContext(ctx, query, id)

	var m models.Mitra
	if err := row.Scan(&m.ID, &m.Nama, &m.Alamat, &m.Kontak, &m.Email, &m.CreatedAt, &m.UpdatedAt); err != nil {
		res.StatusCode = http.StatusNotFound
		res.Message = "Mitra not found"
		return &res, err
	}

	res.StatusCode = http.StatusOK
	res.Message = "Success"
	res.Data = m
	return &res, nil
}

func (s *MitraService) Update(ctx context.Context, data *models.MitraUpdate) (*models.Response, error) {
	var res models.Response
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to begin transaction"
		return &res, err
	}

	query := `UPDATE mitra SET nama = ?, alamat = ?, kontak = ?, email = ?, updated_at = ? WHERE id = ?`
	result, err := tx.ExecContext(ctx, query, data.Nama, data.Alamat, data.Kontak, data.Email, time.Now(), data.ID)
	if err != nil {
		tx.Rollback()
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to update mitra"
		return &res, err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		tx.Rollback()
		res.StatusCode = http.StatusNotFound
		res.Message = "Mitra not found"
		return &res, nil
	}

	tx.Commit()
	res.StatusCode = http.StatusOK
	res.Message = "Mitra updated successfully"
	res.Data = rows
	return &res, nil
}

func (s *MitraService) Delete(ctx context.Context, id uint) (*models.Response, error) {
	var res models.Response
	query := `DELETE FROM mitra WHERE id = ?`
	_, err := s.DB.ExecContext(ctx, query, id)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to delete mitra"
		return &res, err
	}
	res.StatusCode = http.StatusOK
	res.Message = "Mitra deleted successfully"
	return &res, nil
}

func (s *MitraService) GetAll(ctx context.Context, filter *models.MitraFilter) (*models.Response, error) {
	var res models.Response
	var total int64
	mitras := []*models.Mitra{}

	query := `SELECT * FROM mitra`
	countQuery := `SELECT COUNT(*) FROM mitra`

	var filters []string
	var args []interface{}

	if filter.Nama != "" {
		filters = append(filters, "nama LIKE ?")
		args = append(args, "%"+filter.Nama+"%")
	}
	if filter.Email != "" {
		filters = append(filters, "email LIKE ?")
		args = append(args, "%"+filter.Email+"%")
	}
	if filter.Kontak != "" {
		filters = append(filters, "kontak LIKE ?")
		args = append(args, "%"+filter.Kontak+"%")
	}

	if len(filters) > 0 {
		where := " WHERE " + strings.Join(filters, " AND ")
		query += where
		countQuery += where
	}

	if err := s.DB.QueryRowContext(ctx, countQuery, args...).Scan(&total); err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to count records"
		return &res, err
	}

	if filter.PageSize == 0 {
		filter.PageSize = 10
	}
	if filter.PageNumber == 0 {
		filter.PageNumber = 1
	}
	offset := (filter.PageNumber - 1) * filter.PageSize
	query += " LIMIT ? OFFSET ?"
	args = append(args, filter.PageSize, offset)

	rows, err := s.DB.QueryContext(ctx, query, args...)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to query mitra"
		return &res, err
	}
	defer rows.Close()

	for rows.Next() {
		var m models.Mitra
		if err := rows.Scan(&m.ID, &m.Nama, &m.Alamat, &m.Kontak, &m.Email, &m.CreatedAt, &m.UpdatedAt); err != nil {
			res.StatusCode = http.StatusInternalServerError
			res.Message = "Failed to scan mitra"
			return &res, err
		}
		mitras = append(mitras, &m)
	}

	res.StatusCode = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]interface{}{
		"total":  total,
		"result": mitras,
	}
	return &res, nil
}
