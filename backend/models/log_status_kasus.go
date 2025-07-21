package models

import "time"

type LogStatusKasus struct {
	ID            uint      `json:"id"`
	KasusID       uint      `json:"kasus_id"`
	Deskripsi     string    `json:"deskripsi"`        // Misalnya: "Menunggu Ambulans", "Selesai"
	Waktu         time.Time `json:"waktu"`            // Waktu update status
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
