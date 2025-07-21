package models

import "time"

type KategoriKasus struct {
	ID        uint      `json:"id"`
	Nama_kategori      string    `json:"nama_kategori"`
	Deskripsi string    `json:"deskripsi"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
