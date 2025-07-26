package models

import (
	"database/sql"
	"time"
)

type Admin struct {
	ID         uint           `json:"id_admin" validate:"required"`
	Username   string         `json:"username" validate:"required"`
	Password   string         `json:"-"` // Disembunyikan dari JSON
	NamaLengkap string        `json:"nama_lengkap" validate:"required"`
	Email      string         `json:"email" validate:"required,email"`
	NoTelepon  sql.NullString `json:"no_telepon" validate:"required"`
	IDRole     uint           `json:"id_role" validate:"required"`
	CreatedAt  time.Time      `json:"created_at" validate:"required"`
	UpdatedAt  time.Time      `json:"updated_at" validate:"required"`
	DeletedAt  *time.Time      `json:"deleted_at,omitempty"` // Tambahkan field deleted_at
}

type AdminCreate struct {
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	NamaLengkap string `json:"nama_lengkap" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	NoTelepon   string `json:"no_telepon" validate:"required"`
	IDRole      uint   `json:"id_role" validate:"required"`
}


type AdminFilter struct {
	PaginationFilter
	Username   string `json:"username,omitempty" validate:"omitempty"`
	NamaLengkap string `json:"nama_lengkap,omitempty" validate:"omitempty"`
	Email      string `json:"email,omitempty" validate:"omitempty,email"`
	NoTelepon  string `json:"no_telepon,omitempty" validate:"omitempty"`
	IDRole     uint   `json:"id_role,omitempty" validate:"omitempty"`
}