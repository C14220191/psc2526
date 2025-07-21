package models

import (
	"database/sql"
	"time"
)

type Admin struct {
	ID         uint           `json:"id_admin"`
	Username   string         `json:"username"`
	Password   string         `json:"-"` // Disembunyikan dari JSON
	NamaLengkap string        `json:"nama_lengkap"`
	Email      string         `json:"email"`
	NoTelepon  sql.NullString `json:"no_telepon"`
	IDRole     uint           `json:"id_role"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}
