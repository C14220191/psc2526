package models

import (
	"database/sql"
	"time"
)

type FasilitasKesehatan struct {
	ID        uint           `json:"id_faskes"`
	Nama      string         `json:"nama"`
	Tipe      string         `json:"tipe"`       // RSUD, Klinik, Puskesmas, dll
	Alamat    string         `json:"alamat"`
	JamBuka   sql.NullTime   `json:"jam_buka"`
	JamTutup  sql.NullTime   `json:"jam_tutup"`
	Kota      string         `json:"kota"`
	Kontak    string         `json:"kontak"`
	Status    string         `json:"status"`     // aktif / non-aktif
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt *time.Time     `json:"deleted_at,omitempty"`
}

type FasilitasKesehatanCreate struct {
	Nama     string `json:"nama" validate:"required"`
	Tipe     string `json:"tipe" validate:"required"`
	Alamat   string `json:"alamat" validate:"required"`
	JamBuka  string `json:"jam_buka"`   // format: "15:04"
	JamTutup string `json:"jam_tutup"`  // format: "15:04"
	Kota     string `json:"kota" validate:"required"`
	Kontak   string `json:"kontak" validate:"required"`
	Status   string `json:"status" validate:"required"`
}

type FasilitasKesehatanUpdate struct {
	ID       uint   `json:"id" validate:"required"`
	Nama     string `json:"nama" validate:"required"`
	Tipe     string `json:"tipe" validate:"required"`
	Alamat   string `json:"alamat" validate:"required"`
	JamBuka  string `json:"jam_buka"`
	JamTutup string `json:"jam_tutup"`
	Kota     string `json:"kota" validate:"required"`
	Kontak   string `json:"kontak" validate:"required"`
	Status   string `json:"status" validate:"required"`
}

type FasilitasKesehatanFilter struct {
	PaginationFilter
	Nama   string `json:"nama,omitempty"`
	Tipe   string `json:"tipe,omitempty"`
	Kota   string `json:"kota,omitempty"`
	Status string `json:"status,omitempty"`
}
