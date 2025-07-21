package services

import (
	"database/sql"
	"backend/models"
)

type AssessmentService struct {
	DB *sql.DB
}

func NewAssessmentService(db *sql.DB) *AssessmentService {
	return &AssessmentService{DB: db}
}

func (s *AssessmentService) Create(data *models.Assessment) error {
	query := `INSERT INTO assessment (kasus_id, jawaban, created_at, updated_at) VALUES (?, ?, ?, ?)`
	_, err := s.DB.Exec(query, data.KasusID, data.Jawaban, data.CreatedAt, data.UpdatedAt)
	return err
}

func (s *AssessmentService) GetByID(id uint) (*models.Assessment, error) {
	query := `SELECT id, kasus_id, jawaban, created_at, updated_at FROM assessment WHERE id = ?`
	row := s.DB.QueryRow(query, id)

	var result models.Assessment
	err := row.Scan(&result.ID, &result.KasusID, &result.Jawaban, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *AssessmentService) Update(data *models.Assessment) error {
	query := `UPDATE assessment SET kasus_id = ?, jawaban = ?, updated_at = ? WHERE id = ?`
	_, err := s.DB.Exec(query, data.KasusID, data.Jawaban, data.UpdatedAt, data.ID)
	return err
}

func (s *AssessmentService) Delete(id uint) error {
	query := `DELETE FROM assessment WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
