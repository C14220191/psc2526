package models

import "time"

type PenanggungJawabMitra struct {
	ID        uint      `json:"id"`
	MitraID   uint      `json:"mitra_id"`    // FK ke mitra
	Nama      string    `json:"nama"`
	Jabatan   string    `json:"jabatan"`
	Kontak    string    `json:"kontak"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
