package models

import (
    "time"
)

type Assessment struct {
	ID        uint      `json:"id"`
	KasusID   uint      `json:"kasus_id"`
	Jawaban   string    `json:"jawaban"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
