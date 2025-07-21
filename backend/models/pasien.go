package models

import "time"

type Pasien struct {
	ID          uint      `json:"id"`
	Nama        string    `json:"nama"`
	Umur        int       `json:"umur"`
	JenisKelamin string   `json:"jenis_kelamin"`
	TempatLahir string    `json:"tempat_lahir"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	Alamat      string    `json:"alamat"`
	Kontak      string    `json:"kontak"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
