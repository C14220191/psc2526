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

type FasilitasKesehatanService struct {
	DB *sql.DB
}

var _ interfaces.FasilitasKesehatanInterface = &FasilitasKesehatanService{}

func NewFasilitasKesehatanService(db *sql.DB) *FasilitasKesehatanService {
	return &FasilitasKesehatanService{DB: db}
}

func (s *FasilitasKesehatanService) Create(ctx context.Context, data *models.FasilitasKesehatanCreate) (*models.Response, error) {
	var res models.Response
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to begin transaction"
		return &res, err
	}

	query := `
	INSERT INTO fasilitas_kesehatan 
	(nama, tipe, alamat, jam_buka, jam_tutup, kota, kontak, status, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`

	_, err = tx.ExecContext(ctx, query,
		data.Nama, data.Tipe, data.Alamat,
		data.JamBuka, data.JamTutup, data.Kota,
		data.Kontak, data.Status,
	)
	if err != nil {
		tx.Rollback()
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to create fasilitas kesehatan"
		return &res, err
	}
	tx.Commit()
	res.StatusCode = http.StatusCreated
	res.Message = "Fasilitas kesehatan created successfully"
	return &res, nil
}

func (s *FasilitasKesehatanService) GetAll(ctx context.Context, filter *models.FasilitasKesehatanFilter) (*models.Response, error) {
	var res models.Response
	var total int64
	arr := []*models.FasilitasKesehatan{}

	baseQuery := `SELECT id, nama, tipe, alamat, jam_buka, jam_tutup, kota, kontak, status, created_at, updated_at, deleted_at FROM fasilitas_kesehatan`
	countQuery := `SELECT COUNT(*) FROM fasilitas_kesehatan`
	var conditions []string
	var args []any

	conditions = append(conditions, "deleted_at IS NULL")

	if filter.Nama != "" {
		conditions = append(conditions, "nama LIKE ?")
		args = append(args, "%"+filter.Nama+"%")
	}
	if filter.Tipe != "" {
		conditions = append(conditions, "tipe = ?")
		args = append(args, filter.Tipe)
	}
	if filter.Kota != "" {
		conditions = append(conditions, "kota = ?")
		args = append(args, filter.Kota)
	}
	if filter.Status != "" {
		conditions = append(conditions, "status = ?")
		args = append(args, filter.Status)
	}

	if len(conditions) > 0 {
		whereClause := " WHERE " + strings.Join(conditions, " AND ")
		baseQuery += whereClause
		countQuery += whereClause
	}

	if filter.PageSize == 0 {
		filter.PageSize = 10
	}
	if filter.PageNumber == 0 {
		filter.PageNumber = 1
	}

	offset := (filter.PageNumber - 1) * filter.PageSize
	baseQuery += " LIMIT ? OFFSET ?"
	argsWithLimit := append(args, filter.PageSize, offset)

	err := s.DB.QueryRowContext(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to count rows"
		return &res, err
	}

	rows, err := s.DB.QueryContext(ctx, baseQuery, argsWithLimit...)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Query failed"
		return &res, err
	}
	defer rows.Close()

	for rows.Next() {
		var f models.FasilitasKesehatan
		if err := rows.Scan(
			&f.ID, &f.Nama, &f.Tipe, &f.Alamat,
			&f.JamBuka, &f.JamTutup, &f.Kota, &f.Kontak,
			&f.Status, &f.CreatedAt, &f.UpdatedAt, &f.DeletedAt,
		); err != nil {
			res.StatusCode = http.StatusInternalServerError
			res.Message = "Scan error"
			return &res, err
		}
		arr = append(arr, &f)
	}

	res.StatusCode = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]any{
		"total":  total,
		"result": arr,
	}
	return &res, nil
}

func (s *FasilitasKesehatanService) GetByID(ctx context.Context, fasilitas *models.FasilitasKesehatan, id uint) (*models.Response, error) {
	var res models.Response

	query := `SELECT id, nama, tipe, alamat, jam_buka, jam_tutup, kota, kontak, status, created_at, updated_at, deleted_at 
	          FROM fasilitas_kesehatan 
	          WHERE id = ? AND deleted_at IS NULL`

	row := s.DB.QueryRowContext(ctx, query, id)

	var jamBukaStr sql.NullString
	var jamTutupStr sql.NullString
	var deletedAt sql.NullTime

	err := row.Scan(
		&fasilitas.ID, &fasilitas.Nama, &fasilitas.Tipe, &fasilitas.Alamat,
		&jamBukaStr, &jamTutupStr, &fasilitas.Kota, &fasilitas.Kontak,
		&fasilitas.Status, &fasilitas.CreatedAt, &fasilitas.UpdatedAt, &deletedAt,
	)

	if err == sql.ErrNoRows {
		res.StatusCode = http.StatusNotFound
		res.Message = "Data not found"
		return &res, nil
	} else if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Scan error"
		return &res, err
	}

	// Parse jam buka & jam tutup dari string ke time
	if jamBukaStr.Valid {
		t, err := parseTime(jamBukaStr.String)
		if err == nil {
			fasilitas.JamBuka = sql.NullTime{Time: t, Valid: true}
		}
	}
	if jamTutupStr.Valid {
		t, err := parseTime(jamTutupStr.String)
		if err == nil {
			fasilitas.JamTutup = sql.NullTime{Time: t, Valid: true}
		}
	}

	if deletedAt.Valid {
		fasilitas.DeletedAt = &deletedAt.Time
	}

	res.StatusCode = http.StatusOK
	res.Message = "Success"
	res.Data = fasilitas
	return &res, nil
}

// Tambahkan helper ini di bawah fungsi GetByID atau tempat lain
func parseTime(s string) (time.Time, error) {
	// MySQL TIME format: "15:04:05"
	return time.Parse("15:04:05", s)
}


func (s *FasilitasKesehatanService) Update(ctx context.Context, data *models.FasilitasKesehatanUpdate) (*models.Response, error) {
	var res models.Response
	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to begin transaction"
		return &res, err
	}

	query := `UPDATE fasilitas_kesehatan SET nama=?, tipe=?, alamat=?, jam_buka=?, jam_tutup=?, kota=?, kontak=?, status=?, updated_at=NOW() WHERE id=?`
	_, err = tx.ExecContext(ctx, query,
		data.Nama, data.Tipe, data.Alamat,
		data.JamBuka, data.JamTutup, data.Kota,
		data.Kontak, data.Status, data.ID,
	)
	if err != nil {
		tx.Rollback()
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to update fasilitas kesehatan"
		return &res, err
	}
	tx.Commit()
	res.StatusCode = http.StatusOK
	res.Message = "Updated successfully"
	return &res, nil
}


func (s *FasilitasKesehatanService) Delete(id uint) error {
	query := `UPDATE fasilitas_kesehatan SET deleted_at = NOW() WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
