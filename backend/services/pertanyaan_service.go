package services

import (
	"database/sql"
	"backend/models"
)

type PertanyaanService struct {
	DB *sql.DB
}

func NewPertanyaanService(db *sql.DB) *PertanyaanService {
	return &PertanyaanService{DB: db}
}

func (s *PertanyaanService) Create(data *models.Pertanyaan) error {
	query := `INSERT INTO pertanyaan (id_pertanyaan, isi_pertanyaan, jenis, label)
	          VALUES (?, ?, ?, ?)`
	_, err := s.DB.Exec(query,
		data.IDPertanyaan, data.IsiPertanyaan, data.Jenis, data.Label)
	return err
}

func (s *PertanyaanService) GetByID(id int64) (*models.Pertanyaan, error) {
	query := `SELECT id, id_pertanyaan, isi_pertanyaan, jenis, label FROM pertanyaan WHERE id = ?`
	row := s.DB.QueryRow(query, id)

	var result models.Pertanyaan
	err := row.Scan(&result.ID, &result.IDPertanyaan, &result.IsiPertanyaan, &result.Jenis, &result.Label)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *PertanyaanService) Update(data *models.Pertanyaan) error {
	query := `UPDATE pertanyaan 
	          SET id_pertanyaan = ?, isi_pertanyaan = ?, jenis = ?, label = ? 
	          WHERE id = ?`
	_, err := s.DB.Exec(query,
		data.IDPertanyaan, data.IsiPertanyaan, data.Jenis, data.Label, data.ID)
	return err
}

func (s *PertanyaanService) Delete(id int64) error {
	query := `DELETE FROM pertanyaan WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
