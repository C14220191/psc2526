package models

import "time"

type DokumentasiMitra struct {
	ID         uint      `json:"id"`
	IDMitra    uint      `json:"id_mitra" validate:"required"`
	FileURL    string    `json:"file_url" validate:"required,url"`
	Keterangan string    `json:"keterangan" validate:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

