package models

import "time"

type Lokasi struct {
	ID        uint      `json:"id"`
	NamaLokasi string    `json:"nama_lokasi"`
	Alamat     string    `json:"alamat"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	Tipe       string    `json:"tipe"` // kantor, mitra, kasus, rumah sakit, dll
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
