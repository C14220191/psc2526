package models

import "time"

type VerifikasiUser struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	WaktuPengajuan time.Time `json:"waktu_pengajuan"`
	Status    string    `json:"status"` // pending, diterima, ditolak
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
