package services

import (
	"database/sql"
	"backend/models"
)

type LaporanKondisiKorbanService struct {
	DB *sql.DB
}

func NewLaporanKondisiKorbanService(db *sql.DB) *LaporanKondisiKorbanService {
	return &LaporanKondisiKorbanService{DB: db}
}

func (s *LaporanKondisiKorbanService) Create(data *models.LaporanKondisiKorban) error {
	query := `
	INSERT INTO laporan_kondisi_korban (
		kasus_id, kondisi, gejala, tekanan_darah, nadi, suhu_tubuh, kesadaran, deskripsi,
		waktu_pemeriksaan, dilaporkan_oleh, created_at, updated_at
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := s.DB.Exec(query,
		data.KasusID, data.Kondisi, data.Gejala, data.TekananDarah,
		data.Nadi, data.SuhuTubuh, data.Kesadaran, data.Deskripsi,
		data.WaktuPemeriksaan, data.DilaporkanOleh, data.CreatedAt, data.UpdatedAt)
	return err
}

func (s *LaporanKondisiKorbanService) GetByID(id uint) (*models.LaporanKondisiKorban, error) {
	query := `
	SELECT id, kasus_id, kondisi, gejala, tekanan_darah, nadi, suhu_tubuh, kesadaran, deskripsi,
	waktu_pemeriksaan, dilaporkan_oleh, created_at, updated_at
	FROM laporan_kondisi_korban WHERE id = ?`
	row := s.DB.QueryRow(query, id)

	var result models.LaporanKondisiKorban
	err := row.Scan(
		&result.ID, &result.KasusID, &result.Kondisi, &result.Gejala, &result.TekananDarah,
		&result.Nadi, &result.SuhuTubuh, &result.Kesadaran, &result.Deskripsi,
		&result.WaktuPemeriksaan, &result.DilaporkanOleh, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *LaporanKondisiKorbanService) Update(data *models.LaporanKondisiKorban) error {
	query := `
	UPDATE laporan_kondisi_korban SET
	kondisi = ?, gejala = ?, tekanan_darah = ?, nadi = ?, suhu_tubuh = ?, kesadaran = ?, 
	deskripsi = ?, waktu_pemeriksaan = ?, dilaporkan_oleh = ?, updated_at = ?
	WHERE id = ?`
	_, err := s.DB.Exec(query,
		data.Kondisi, data.Gejala, data.TekananDarah, data.Nadi,
		data.SuhuTubuh, data.Kesadaran, data.Deskripsi,
		data.WaktuPemeriksaan, data.DilaporkanOleh, data.UpdatedAt,
		data.ID)
	return err
}

func (s *LaporanKondisiKorbanService) Delete(id uint) error {
	query := `DELETE FROM laporan_kondisi_korban WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
