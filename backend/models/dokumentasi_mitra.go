package models

import "time"

type DokumentasiMitra struct {
	ID         uint      `json:"id"`
	IDMitra    uint      `json:"id_mitra"`     // FK ke mitra
	FileURL    string    `json:"file_url"`     // URL file dokumentasi
	Keterangan string    `json:"keterangan"`   // Deskripsi / caption
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
