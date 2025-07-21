package models

import "time"

type Dokter struct {
	ID            uint      `json:"id"`
	Nama          string    `json:"nama"`
	Spesialisasi  string    `json:"spesialisasi"`   // Umum, Bedah, dll
	FaskesID      uint      `json:"faskes_id"`       // FK ke FasilitasKesehatan
	NoSTR         string    `json:"no_str"`          // Nomor STR
	JenisKelamin  string    `json:"jenis_kelamin"`   // Laki-laki, Perempua
	Kontak        string    `json:"kontak"`          // No HP
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
