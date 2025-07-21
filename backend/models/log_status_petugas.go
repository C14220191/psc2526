package models

import "time"

type LogStatusPetugas struct {
	ID         uint      `json:"id"`
	PetugasID  uint      `json:"petugas_id"`
	Status     string    `json:"status"`     // aktif, tidak tersedia, dll
	Waktu	 time.Time `json:"waktu"`      // Waktu perubahan status
	Timestamp  time.Time `json:"timestamp"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

