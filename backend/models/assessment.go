package models

import (
	"time"
)

type Assessment struct {
	ID        uint       `json:"id" validate:"required"`
	KasusID   uint       `json:"kasus_id" validate:"required"`
	Jawaban   string     `json:"jawaban" validate:"required"`
	CreatedAt time.Time  `json:"created_at" validate:"required"`
	UpdatedAt time.Time  `json:"updated_at" validate:"required"`
	DeletedAt *time.Time `json:"deleted_at" validate:"omitempty"` // Tambahkan field deleted_at
}

type AssessmentCreate struct {
	KasusID uint   `json:"kasus_id" validate:"required"`
	Jawaban string `json:"jawaban" validate:"required"`
}

type AssessmentUpdate struct {
	ID      uint   `json:"id" validate:"required"`
	KasusID uint   `json:"kasus_id" validate:"required"`
	Jawaban string `json:"jawaban" validate:"required"`
}

type AssessmentGetAllResponse struct {
	PaginationFilter
	ID        uint      `json:"id" validate:"required"`
	KasusID   uint      `json:"kasus_id" validate:"required"`
	Jawaban   string    `json:"jawaban" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" validate:"required"`
}
