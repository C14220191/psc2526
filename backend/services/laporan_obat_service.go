package services

import (
	"database/sql"
	"backend/models"
)

type LaporanObatService struct {
	DB *sql.DB
}

func NewLaporanObatService(db *sql.DB) *LaporanObatService {
	return &LaporanObatService{DB: db}
}

func (s *LaporanObatService) Create(data *models.LaporanObat) error {
	query := `
		INSERT INTO laporan_obat (
			kasus_id, nama_obat, dosis, jumlah,
			cara_pemberian, waktu_pemberian, keterangan,
			petugas_id, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := s.DB.Exec(query,
		data.KasusID, data.NamaObat, data.Dosis, data.Jumlah,
		data.CaraPemberian, data.WaktuPemberian, data.Keterangan,
		data.PetugasID, data.CreatedAt, data.UpdatedAt,
	)
	return err
}

func (s *LaporanObatService) GetByID(id uint) (*models.LaporanObat, error) {
	query := `
		SELECT id, kasus_id, nama_obat, dosis, jumlah,
		       cara_pemberian, waktu_pemberian, keterangan,
		       petugas_id, created_at, updated_at
		FROM laporan_obat WHERE id = ?`
	row := s.DB.QueryRow(query, id)

	var result models.LaporanObat
	err := row.Scan(
		&result.ID, &result.KasusID, &result.NamaObat, &result.Dosis, &result.Jumlah,
		&result.CaraPemberian, &result.WaktuPemberian, &result.Keterangan,
		&result.PetugasID, &result.CreatedAt, &result.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *LaporanObatService) Update(data *models.LaporanObat) error {
	query := `
		UPDATE laporan_obat SET
			kasus_id = ?, nama_obat = ?, dosis = ?, jumlah = ?,
			cara_pemberian = ?, waktu_pemberian = ?, keterangan = ?,
			petugas_id = ?, updated_at = ?
		WHERE id = ?`
	_, err := s.DB.Exec(query,
		data.KasusID, data.NamaObat, data.Dosis, data.Jumlah,
		data.CaraPemberian, data.WaktuPemberian, data.Keterangan,
		data.PetugasID, data.UpdatedAt, data.ID,
	)
	return err
}

func (s *LaporanObatService) Delete(id uint) error {
	query := `DELETE FROM laporan_obat WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
