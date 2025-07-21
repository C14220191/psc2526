package services

import (
	"database/sql"
	"backend/models"
)

type UserService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) Create(data *models.User) error {
	query := `INSERT INTO user 
	(username, password, nama_lengkap, nik, no_telepon, alamat, kota, jenis_kelamin, tanggal_lahir, status, otp_status, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := s.DB.Exec(query,
		data.Username, data.Password, data.NamaLengkap, data.NIK, data.NoTelepon,
		data.Alamat, data.Kota, data.JenisKelamin, data.TanggalLahir, data.Status,
		data.OTPStatus, data.CreatedAt, data.UpdatedAt,
	)
	return err
}

func (s *UserService) GetByID(id uint) (*models.User, error) {
	query := `SELECT id, username, password, nama_lengkap, nik, no_telepon, alamat, kota, jenis_kelamin, tanggal_lahir, status, otp_status, created_at, updated_at FROM user WHERE id = ?`
	row := s.DB.QueryRow(query, id)
	var result models.User
	err := row.Scan(
		&result.ID, &result.Username, &result.Password, &result.NamaLengkap, &result.NIK,
		&result.NoTelepon, &result.Alamat, &result.Kota, &result.JenisKelamin, &result.TanggalLahir,
		&result.Status, &result.OTPStatus, &result.CreatedAt, &result.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *UserService) Update(data *models.User) error {
	query := `UPDATE user SET 
	username = ?, password = ?, nama_lengkap = ?, nik = ?, no_telepon = ?, alamat = ?, kota = ?, jenis_kelamin = ?, tanggal_lahir = ?, status = ?, otp_status = ?, updated_at = ?
	WHERE id = ?`
	_, err := s.DB.Exec(query,
		data.Username, data.Password, data.NamaLengkap, data.NIK, data.NoTelepon,
		data.Alamat, data.Kota, data.JenisKelamin, data.TanggalLahir, data.Status,
		data.OTPStatus, data.UpdatedAt, data.ID,
	)
	return err
}

func (s *UserService) Delete(id uint) error {
	query := `DELETE FROM user WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
