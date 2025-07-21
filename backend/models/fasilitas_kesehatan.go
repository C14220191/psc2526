package models

import "time"

type FasilitasKesehatan struct {
	ID        uint      `json:"id"`
	Nama      string    `json:"nama"`       // Nama RS / Klinik / Puskesmas
	Tipe     string    `json:"tipe"`      // RSUD, Puskesmas, dll
	Alamat    string    `json:"alamat"`     // Alamat lengkap
	JamBuka   time.Time    `json:"jam_buka"`   // Jam buka operasional
	JamTutup  time.Time    `json:"jam_tutup"`  // Jam tutup
	Kota      string    `json:"kota"`       // Kota atau kabupaten
	Kontak    string    `json:"kontak"`     // No. telepon
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
