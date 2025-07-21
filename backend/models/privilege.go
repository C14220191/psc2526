package models

import "time"

type Privilege struct {
	ID uint `json:"id"`
	NamaPrivilege string    `json:"nama_privilege"` // Nama privilege, misal "admin", "petugas", "mitra"
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"` // Waktu penghapusan, jika ada
}
