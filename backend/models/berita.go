package models

import (
	"time"
)

type Berita struct {
	ID        uint       `json:"id_berita"`
	IDAdmin   uint       `json:"id_admin`
	Judul     string     `json:"judul"`
	Isi       string     `json:"isi"`
	Thumbnail string     `json:"thumbnail"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type BeritaCreate struct {
	IDAdmin   uint   `json:"id_admin" validate:"required"`
	Judul     string `json:"judul" validate:"required"`
	Isi       string `json:"isi" validate:"required"`
	Thumbnail string `json:"thumbnail" validate:"filepath"`
}

type BeritaUpdate struct {
	IDAdmin   uint   `json:"id_admin" validate:"required"`
	Judul     string `json:"judul" validate:"required"`
	Isi       string `json:"isi" validate:"required"`
	Thumbnail string `json:"thumbnail" validate:"omitempty"`
}

type BeritaGet struct {
	PaginationFilter
	ID        uint   `json:"id_berita,omitempty"`
	IDAdmin   uint   `json:"id_admin,omitempty"`
	Judul     string `json:"judul,omitempty"`
	Isi       string `json:"isi,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}
