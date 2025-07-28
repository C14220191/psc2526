package services

import (
	"backend/interfaces"
	"backend/models"
	"context"
	"database/sql"
	"net/http"
)

type LaporanKondisiKorbanService struct {
	DB *sql.DB
}

var _ interfaces.LaporanKondisiKorbanInterface = &LaporanKondisiKorbanService{}

func NewLaporanKondisiKorbanService(db *sql.DB) *LaporanKondisiKorbanService {
	return &LaporanKondisiKorbanService{DB: db}
}

func (s *LaporanKondisiKorbanService) Create(ctx context.Context, data *models.LaporanKondisiKorban) (*models.Response, error) {
	var res models.Response
	query := `
	INSERT INTO laporan_kondisi_korban (
		kasus_id, kondisi, gejala, tekanan_darah, nadi, suhu_tubuh, kesadaran, deskripsi,
		waktu_pemeriksaan, dilaporkan_oleh, created_at, updated_at
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := s.DB.ExecContext(ctx, query,
		data.KasusID, data.Kondisi, data.Gejala, data.TekananDarah,
		data.Nadi, data.SuhuTubuh, data.Kesadaran, data.Deskripsi,
		data.WaktuPemeriksaan, data.DilaporkanOleh, data.CreatedAt, data.UpdatedAt,
	)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to create laporan"
		return &res, err
	}
	res.StatusCode = http.StatusCreated
	res.Message = "Laporan kondisi korban created successfully"
	return &res, nil
}

func (s *LaporanKondisiKorbanService) GetByID(ctx context.Context, laporan *models.LaporanKondisiKorban, id uint) (*models.Response, error) {
	var res models.Response
	query := `SELECT id, kasus_id, kondisi, gejala, tekanan_darah, nadi, suhu_tubuh, kesadaran, deskripsi,
	waktu_pemeriksaan, dilaporkan_oleh, created_at, updated_at
	FROM laporan_kondisi_korban WHERE id = ?`

	row := s.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&laporan.ID, &laporan.KasusID, &laporan.Kondisi, &laporan.Gejala, &laporan.TekananDarah,
		&laporan.Nadi, &laporan.SuhuTubuh, &laporan.Kesadaran, &laporan.Deskripsi,
		&laporan.WaktuPemeriksaan, &laporan.DilaporkanOleh, &laporan.CreatedAt, &laporan.UpdatedAt,
	)
	if err != nil {
		res.StatusCode = http.StatusNotFound
		res.Message = "Laporan not found"
		return &res, err
	}

	res.StatusCode = http.StatusOK
	res.Message = "Success"
	res.Data = laporan
	return &res, nil
}

func (s *LaporanKondisiKorbanService) Update(ctx context.Context, data *models.LaporanKondisiKorban) (*models.Response, error) {
	var res models.Response
	query := `
	UPDATE laporan_kondisi_korban SET
	kondisi = ?, gejala = ?, tekanan_darah = ?, nadi = ?, suhu_tubuh = ?, kesadaran = ?, 
	deskripsi = ?, waktu_pemeriksaan = ?, dilaporkan_oleh = ?, updated_at = ?
	WHERE id = ?`

	_, err := s.DB.ExecContext(ctx, query,
		data.Kondisi, data.Gejala, data.TekananDarah, data.Nadi,
		data.SuhuTubuh, data.Kesadaran, data.Deskripsi,
		data.WaktuPemeriksaan, data.DilaporkanOleh, data.UpdatedAt,
		data.ID,
	)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to update laporan"
		return &res, err
	}
	res.StatusCode = http.StatusOK
	res.Message = "Laporan updated successfully"
	return &res, nil
}

func (s *LaporanKondisiKorbanService) Delete(id uint) error {
	query := `DELETE FROM laporan_kondisi_korban WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
