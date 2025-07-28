package models

import (
	"database/sql"
	"time"
)

type Dokter struct {
	ID           uint           `json:"id_dokter"`
	Nama         string         `json:"nama"`
	Bidang       sql.NullString `json:"bidang"`
	NoHP         sql.NullString `json:"no_hp"`
	JenisKelamin sql.NullString `json:"jenis_kelamin"`
	IDFaskes     uint           `json:"id_faskes"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    *time.Time     `json:"deleted_at,omitempty"`
}

type DokterCreate struct {
	Nama         string `json:"nama" validate:"required"`
	Bidang       string `json:"bidang"`
	NoHP         string `json:"no_hp"`
	JenisKelamin string `json:"jenis_kelamin"`
	IDFaskes     uint   `json:"id_faskes" validate:"required"`
}

type DokterUpdate struct {
	ID           uint   `json:"id" validate:"required"`
	Nama         string `json:"nama" validate:"required"`
	Bidang       string `json:"bidang"`
	NoHP         string `json:"no_hp"`
	JenisKelamin string `json:"jenis_kelamin"`
	IDFaskes     uint   `json:"id_faskes" validate:"required"`
}

type DokterFilter struct {
	PaginationFilter
	Nama         string `json:"nama,omitempty"`
	Bidang       string `json:"bidang,omitempty"`
	NoHP         string `json:"no_hp,omitempty"`
	JenisKelamin string `json:"jenis_kelamin,omitempty"`
	IDFaskes     uint   `json:"id_faskes,omitempty"`
}
