package services

import (
	"backend/interfaces"
	"backend/models"
	"context"
	"database/sql"
	"fmt"
	"net/http"
)

type AssessmentService struct {
	DB *sql.DB
}

func NewAssessmentService(db *sql.DB) *AssessmentService {
	return &AssessmentService{DB: db}
}

var _ interfaces.AssessmentInterface = &AssessmentService{}

func (s *AssessmentService) Create(data *models.AssessmentCreate, ctx context.Context) (*models.Response, error) {
	var res models.Response

	tx, err := s.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to begin transaction"
		res.Data = nil
		return &res, err
	}
	query := `INSERT INTO assessment (kasus_id, jawaban) VALUES (?, ?)`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		tx.Rollback()
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to prepare statement"
		res.Data = nil
		return &res, err
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(ctx, data.KasusID, data.Jawaban)
	if err != nil {
		tx.Rollback()
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to execute query"
		res.Data = nil
		return &res, err
	}
	lastID, _ := result.LastInsertId()
	if err := tx.Commit(); err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to commit transaction"
		res.Data = nil
		return &res, err
	}
	res.StatusCode = http.StatusCreated
	res.Message = "Assessment created successfully"
	res.Data = lastID

	return &res, nil
}
func (s *AssessmentService) GetAll(ctx context.Context, data models.AssessmentGetAllResponse) (*models.Response, error) {
	var res models.Response
	fmt.Printf("PageNumber: %d, PageSize: %d\n", data.PageNumber, data.PageSize)
	if data.PageNumber == 0 {
		data.PageNumber = 1
	}
	if data.PageSize == 0 {
		data.PageSize = 10
	}

	offset := (data.PageNumber - 1) * data.PageSize
	query := `SELECT id, kasus_id, jawaban, created_at, updated_at FROM assessment WHERE deleted_at IS NULL LIMIT ? OFFSET ?`
	rows, err := s.DB.QueryContext(ctx, query, data.PageSize, offset)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to execute query"
		res.Data = nil
		return &res, err
	}

	defer rows.Close()
	var arr []models.Assessment
	for rows.Next() {
		var assessment models.Assessment
		if err := rows.Scan(&assessment.ID, &assessment.KasusID, &assessment.Jawaban, &assessment.CreatedAt, &assessment.UpdatedAt); err != nil {
			res.StatusCode = http.StatusInternalServerError
			res.Message = "Failed to scan row"
			res.Data = nil
			return &res, err
		}
		arr = append(arr, assessment)
	}
	fmt.Printf("PageNumber: %d, PageSize: %d, Offset: %d", data.PageNumber, data.PageSize, offset)
	fmt.Printf("Jumlah data hasil query: %d", len(arr))
	res.StatusCode = http.StatusOK
	res.Message = "Assessments retrieved successfully"
	res.Data = arr
	return &res, nil
}

func (s *AssessmentService) GetByID(ctx context.Context, assessment models.Assessment, id uint) (*models.Response, error) {
	var res models.Response

	query := `SELECT id, kasus_id, jawaban, created_at, updated_at FROM assessment WHERE id = ?`
	row := s.DB.QueryRowContext(ctx, query, id)
	if row.Err() != nil {
		res.StatusCode = http.StatusNotFound
		res.Message = "Assessment not found"
		res.Data = nil
		return &res, row.Err()
	}

	var result models.Assessment
	err := row.Scan(&result.ID, &result.KasusID, &result.Jawaban, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to scan row"
		res.Data = nil
		return &res, err
	}
	res.StatusCode = http.StatusOK
	res.Message = "Assessment retrieved successfully"
	res.Data = result
	return &res, nil
}

func (s *AssessmentService) Update(data *models.AssessmentUpdate, ctx context.Context) (*models.Response, error) {
	var res models.Response
	tx, err := s.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to begin transaction"
		res.Data = nil
		return &res, err
	}

	query := `UPDATE assessment SET kasus_id = ?, jawaban = ? WHERE id = ?`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		tx.Rollback()
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to prepare statement"
		res.Data = nil
		return &res, err
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(ctx, data.KasusID, data.Jawaban, data.ID)
	if err != nil {
		tx.Rollback()
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to execute query"
		res.Data = nil
		return &res, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to get rows affected"
		res.Data = nil
		return &res, err
	}
	if rowsAffected == 0 {
		tx.Rollback()
		res.StatusCode = http.StatusNotFound
		res.Message = "No assessment found with the given ID"
		res.Data = nil
		return &res, nil
	}
	
	if err := tx.Commit(); err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to commit transaction"
		res.Data = nil
		return &res, err
	}
	res.StatusCode = http.StatusOK
	res.Message = "Assessment updated successfully"
	res.Data = data.ID
	return &res, nil
}

func (s *AssessmentService) Delete(id uint, ctx context.Context) (*models.Response, error) {
	var res models.Response
	tx, err := s.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to begin transaction"
		res.Data = nil
		return &res, err
	}
	query := `UPDATE assessment SET deleted_at = NOW() WHERE id = ?`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		tx.Rollback()
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to prepare statement"
		res.Data = nil
		return &res, err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		tx.Rollback()
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to execute query"
		res.Data = nil
		return &res, err
	}
	if err := tx.Commit(); err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to commit transaction"
		res.Data = nil
		return &res, err
	}
	res.StatusCode = http.StatusOK
	res.Message = "Assessment deleted successfully"
	res.Data = nil

	return &res, nil
}
