package models

import "time"

type KategoriPelaporan struct {
	ID        uint      `json:"id"`
	Kode      string    `json:"kode"`       // Kode unik seperti MED, KEC, dll
	Nama      string    `json:"nama"`       // Nama kategori
	Deskripsi string    `json:"deskripsi"`  // Penjelasan
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
