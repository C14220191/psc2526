package models

import "time"

// Pelapor merepresentasikan tabel 'pelapor'
type Pelapor struct {
	ID            int64     `json:"id"`              // PK (auto increment)
	IDPelapor     string    `json:"id_pelapor"`      // Unique identifier (UQ)
	NamaPelapor   string    `json:"nama_pelapor"`    // Not Null
	LokasiPelapor string    `json:"lokasi_pelapor"`  // Nullable
	NoTelp        string    `json:"no_telp"`         // String, bisa panjang
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
