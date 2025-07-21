package models

import "time"

type Kendaraan struct {
	ID         uint      `json:"id"`
	NomorPlat  string    `json:"nomor_plat"`  // Unik
	Jenis      string    `json:"jenis"`       // Ambulans, Motor, dll
	Merek      string    `json:"merek"`
	Warna      string    `json:"warna"`
	Tahun      int       `json:"tahun"`
	Status     string    `json:"status"`      // aktif, rusak, standby, dll
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
