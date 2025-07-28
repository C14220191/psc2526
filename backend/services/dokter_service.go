package services

import (
	"backend/interfaces"
	"backend/models"
	"context"
	"database/sql"
	"net/http"
	"strings"
	"time"
)

type DokterService struct {
	DB *sql.DB
}

var _ interfaces.DokterInterface = &DokterService{}

func NewDokterService(db *sql.DB) *DokterService {
	return &DokterService{DB: db}
}

func (s *DokterService) Create(ctx context.Context, data *models.DokterCreate) (*models.Response, error) {
	var res models.Response
	query := `INSERT INTO dokter (nama, bidang, no_hp, jenis_kelamin, id_faskes, created_at, updated_at)
	          VALUES (?, ?, ?, ?, ?, NOW(), NOW())`
	result, err := s.DB.ExecContext(ctx, query, data.Nama, data.Bidang, data.NoHP, data.JenisKelamin, data.IDFaskes)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to create dokter"
		return &res, err
	}
	lastID, _ := result.LastInsertId()
	res.StatusCode = http.StatusCreated
	res.Message = "Dokter created successfully"
	res.Data = lastID
	return &res, nil
}

func (s *DokterService) GetByID(ctx context.Context, dokter *models.Dokter, id uint) (*models.Response, error) {
	var res models.Response
	query := `SELECT id, nama, bidang, no_hp, jenis_kelamin, id_faskes, created_at, updated_at, deleted_at
	          FROM dokter WHERE id = ? AND deleted_at IS NULL`
	row := s.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&dokter.ID, &dokter.Nama, &dokter.Bidang, &dokter.NoHP,
		&dokter.JenisKelamin, &dokter.IDFaskes, &dokter.CreatedAt,
		&dokter.UpdatedAt, &dokter.DeletedAt,
	)
	if err != nil {
		res.StatusCode = http.StatusNotFound
		res.Message = "Dokter not found"
		return &res, err
	}
	res.StatusCode = http.StatusOK
	res.Message = "Success"
	res.Data = dokter
	return &res, nil
}

func (s *DokterService) GetAll(ctx context.Context, filter *models.DokterFilter) (*models.Response, error) {
	var res models.Response
	var total int64
	dokters := []*models.Dokter{}

	query := `SELECT id, nama, bidang, no_hp, jenis_kelamin, id_faskes, created_at, updated_at, deleted_at FROM dokter`
	countQuery := `SELECT COUNT(*) FROM dokter`

	var sqlFilters []string
	var sqlVars []any

	if filter.PageSize == 0 {
		filter.PageSize = 10
	}
	if filter.PageNumber == 0 {
		filter.PageNumber = 1
	}

	sqlFilters = append(sqlFilters, "deleted_at IS NULL")

	if filter.Nama != "" {
		sqlFilters = append(sqlFilters, "nama LIKE ?")
		sqlVars = append(sqlVars, "%"+filter.Nama+"%")
	}
	if filter.Bidang != "" {
		sqlFilters = append(sqlFilters, "bidang LIKE ?")
		sqlVars = append(sqlVars, "%"+filter.Bidang+"%")
	}
	if filter.NoHP != "" {
		sqlFilters = append(sqlFilters, "no_hp LIKE ?")
		sqlVars = append(sqlVars, "%"+filter.NoHP+"%")
	}
	if filter.JenisKelamin != "" {
		sqlFilters = append(sqlFilters, "jenis_kelamin LIKE ?")
		sqlVars = append(sqlVars, "%"+filter.JenisKelamin+"%")
	}
	if filter.IDFaskes != 0 {
		sqlFilters = append(sqlFilters, "id_faskes = ?")
		sqlVars = append(sqlVars, filter.IDFaskes)
	}

	if len(sqlFilters) > 0 {
		condition := " WHERE " + strings.Join(sqlFilters, " AND ")
		query += condition
		countQuery += condition
	}

	row := s.DB.QueryRowContext(ctx, countQuery, sqlVars...)
	if err := row.Scan(&total); err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to count rows"
		return &res, err
	}

	query += " LIMIT ? OFFSET ?"
	offset := (filter.PageNumber - 1) * filter.PageSize
	sqlVars = append(sqlVars, filter.PageSize, offset)

	rows, err := s.DB.QueryContext(ctx, query, sqlVars...)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Query error"
		return &res, err
	}
	defer rows.Close()

	for rows.Next() {
		var d models.Dokter
		if err := rows.Scan(
			&d.ID, &d.Nama, &d.Bidang, &d.NoHP,
			&d.JenisKelamin, &d.IDFaskes,
			&d.CreatedAt, &d.UpdatedAt, &d.DeletedAt,
		); err != nil {
			res.StatusCode = http.StatusInternalServerError
			res.Message = "Scan error"
			return &res, err
		}
		dokters = append(dokters, &d)
	}

	res.StatusCode = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]any{
		"total":  total,
		"result": dokters,
	}
	return &res, nil
}

func (s *DokterService) Update(ctx context.Context, data *models.DokterUpdate) (*models.Response, error) {
	var res models.Response
	query := `UPDATE dokter SET nama = ?, bidang = ?, no_hp = ?, jenis_kelamin = ?, id_faskes = ?, updated_at = ? WHERE id = ?`
	result, err := s.DB.ExecContext(ctx, query,
		data.Nama, data.Bidang, data.NoHP, data.JenisKelamin,
		data.IDFaskes, time.Now(), data.ID,
	)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Update failed"
		return &res, err
	}
	affected, _ := result.RowsAffected()
	res.StatusCode = http.StatusOK
	res.Message = "Dokter updated successfully"
	res.Data = affected
	return &res, nil
}

func (s *DokterService) Delete(id uint) error {
	query := `UPDATE dokter SET deleted_at = NOW() WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
