package services

import (
	"backend/interfaces"
	"backend/models"
	"context"
	"database/sql"
	"net/http"
)

type DokumentasiMitraService struct {
	DB *sql.DB
}

var _ interfaces.DokumentasiMitraInterface = &DokumentasiMitraService{}

func NewDokumentasiMitraService(db *sql.DB) *DokumentasiMitraService {
	return &DokumentasiMitraService{DB: db}
}

func (s *DokumentasiMitraService) Create(ctx context.Context, data *models.DokumentasiMitra) (*models.Response, error) {
	var res models.Response

	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to start transaction"
		return &res, err
	}

	query := `INSERT INTO dokumentasi_mitra (id_mitra, file_url, keterangan, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	_, err = tx.ExecContext(ctx, query, data.IDMitra, data.FileURL, data.Keterangan, data.CreatedAt, data.UpdatedAt)
	if err != nil {
		tx.Rollback()
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to create dokumentasi mitra"
		return &res, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Transaction commit failed"
		return &res, err
	}

	res.StatusCode = http.StatusCreated
	res.Message = "Dokumentasi mitra created successfully"
	return &res, nil
}


func (s *DokumentasiMitraService) GetByID(ctx context.Context, dokumentasi *models.DokumentasiMitra, id uint) (*models.Response, error) {
	var res models.Response
	query := `SELECT id, id_mitra, file_url, keterangan, created_at, updated_at FROM dokumentasi_mitra WHERE id = ?`
	row := s.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&dokumentasi.ID, &dokumentasi.IDMitra, &dokumentasi.FileURL, &dokumentasi.Keterangan, &dokumentasi.CreatedAt, &dokumentasi.UpdatedAt)
	if err != nil {
		res.StatusCode = http.StatusNotFound
		res.Message = "Dokumentasi mitra not found"
		return &res, err
	}
	res.StatusCode = http.StatusOK
	res.Message = "Success"
	res.Data = dokumentasi
	return &res, nil
}

func (s *DokumentasiMitraService) Update(ctx context.Context, data *models.DokumentasiMitra) (*models.Response, error) {
	var res models.Response

	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to start transaction"
		return &res, err
	}

	query := `
	UPDATE dokumentasi_mitra
	SET id_mitra = ?, file_url = ?, keterangan = ?, updated_at = ?
	WHERE id = ?`

	_, err = tx.ExecContext(ctx, query, data.IDMitra, data.FileURL, data.Keterangan, data.UpdatedAt, data.ID)
	if err != nil {
		tx.Rollback()
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to update dokumentasi mitra"
		return &res, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Transaction commit failed"
		return &res, err
	}

	res.StatusCode = http.StatusOK
	res.Message = "Dokumentasi mitra updated successfully"
	return &res, nil
}


func (s *DokumentasiMitraService) Delete(id uint) error {
	query := `DELETE FROM dokumentasi_mitra WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
