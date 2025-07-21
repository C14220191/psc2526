package models

import "time"

type Kasus struct {
	ID                 uint       `json:"id"`
	Judul              string     `json:"judul"`
	Deskripsi          string     `json:"deskripsi"`
	KategoriKasusID    uint       `json:"kategori_kasus_id"`
	KoordinatLatitude  float64    `json:"koordinat_latitude"`
	KoordinatLongitude float64    `json:"koordinat_longitude"`
	AlamatLengkap      string     `json:"alamat_lengkap"`
	Status             string     `json:"status"`
	WaktuKejadian      time.Time  `json:"waktu_kejadian"`
	PelaporID          uint       `json:"pelapor_id"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at,omitempty"`
}
