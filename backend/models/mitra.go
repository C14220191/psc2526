package models

import "time"

type Mitra struct {
	ID        uint      `json:"id"`
	Nama      string    `json:"nama"`
	Alamat    string    `json:"alamat"`
	Kontak    string    `json:"kontak"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
