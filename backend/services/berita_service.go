package services

import (
	"backend/interfaces"
	"backend/models"
	"context"
	"database/sql"
	"net/http"
)

type BeritaService struct {
	DB *sql.DB
}

func NewBeritaService(db *sql.DB) *BeritaService {
	return &BeritaService{DB: db}
}

var _ interfaces.BeritaInterface = &BeritaService{}

func (s *BeritaService) Create(data *models.BeritaCreate, ctx context.Context) (*models.Response, error) {
	var res models.Response
	tx, err := s.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		res.Message = "Failed to begin transaction"
		res.StatusCode = http.StatusInternalServerError
		res.Data = nil
		return &res, err
	}

	query := `
	INSERT INTO berita (id_admin, judul, isi, thumbnail)
	VALUES (?, ?, ?, ?)
	`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		res.Message = "Failed to prepare statement"
		res.StatusCode = http.StatusInternalServerError
		res.Data = nil
		return &res, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx,
		data.IDAdmin,
		data.Judul,
		data.Isi,
		data.Thumbnail,
	)
	if err != nil {
		res.Message = "Failed to execute statement"
		res.StatusCode = http.StatusInternalServerError
		res.Data = nil
		return &res, err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		res.Message = "Failed to get last insert ID"
		res.StatusCode = http.StatusInternalServerError
		res.Data = nil
		return &res, err
	}
	if err := tx.Commit(); err != nil {
		res.Message = "Failed to commit transaction"
		res.StatusCode = http.StatusInternalServerError
		res.Data = nil
		return &res, err
	}
	res.Message = "Berita created successfully"
	res.StatusCode = http.StatusCreated
	res.Data = lastID

	return &res, nil
}

func (s *BeritaService) GetByID(id uint, ctx context.Context) (*models.Response, error) {
	var res models.Response
	query := `
	SELECT id, id_admin, judul, isi, thumbnail FROM berita WHERE id = ?
	`
	stmt, err := s.DB.PrepareContext(ctx, query)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to prepare statement"
		res.Data = nil
		return &res, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)
	if row.Err() != nil {
		res.StatusCode = http.StatusNotFound
		res.Message = "Berita not found"
		res.Data = nil
		return &res, row.Err()
	}
	var result models.BeritaGet
	err = row.Scan(
		&result.ID,
		&result.IDAdmin,
		&result.Judul,
		&result.Isi,
		&result.Thumbnail,
	)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to scan row"
		res.Data = nil
		return &res, err
	}
	res.StatusCode = http.StatusOK
	res.Message = "Berita retrieved successfully"
	res.Data = result

	return &res, nil
}

func (s *BeritaService) Update(data *models.BeritaUpdate, ctx context.Context, id uint) (*models.Response, error) {
	var res models.Response

	tx, err := s.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		res.Message = "Failed to begin transaction"
		res.StatusCode = http.StatusInternalServerError
		res.Data = nil
		return &res, err
	}

	query := `
	UPDATE berita
	SET id_admin = ?, judul = ?, isi = ?, thumbnail = ? 
	WHERE id = ?
	`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		res.Message = "Failed to prepare statement"
		res.StatusCode = http.StatusInternalServerError
		res.Data = nil
		tx.Rollback()
		return &res, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx,
		data.IDAdmin, data.Judul, data.Isi, data.Thumbnail, id)
	if err != nil {
		res.Message = "Failed to execute statement"
		res.StatusCode = http.StatusInternalServerError
		res.Data = nil
		tx.Rollback()
		return &res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		res.Message = "Failed to check affected rows"
		res.StatusCode = http.StatusInternalServerError
		res.Data = nil
		tx.Rollback()
		return &res, err
	}

	if rowsAffected == 0 {
		res.Message = "No berita found with the given ID"
		res.StatusCode = http.StatusNotFound
		res.Data = nil
		tx.Rollback()
		return &res, nil
	}

	if err := tx.Commit(); err != nil {
		res.Message = "Failed to commit transaction"
		res.StatusCode = http.StatusInternalServerError
		res.Data = nil
		return &res, err
	}

	res.Message = "Berita updated successfully"
	res.StatusCode = http.StatusOK
	res.Data = id
	return &res, nil
}

func (s *BeritaService) Delete(id uint, ctx context.Context) (*models.Response, error) {
	var res models.Response
	query := `
	UPDATE berita SET deleted_at = Now() WHERE id = ?
	`
	stmt, err := s.DB.PrepareContext(ctx, query)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to prepare statement"
		res.Data = nil
		return &res, err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to execute query"
		res.Data = nil
		return &res, err
	}
	res.StatusCode = http.StatusOK
	res.Message = "Berita deleted successfully"
	res.Data = nil
	return &res, nil
}
