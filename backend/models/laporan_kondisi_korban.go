package models

import "time"

type LaporanKondisiKorban struct {
	ID            uint      `json:"id"`
	KasusID       uint      `json:"kasus_id"`
	Kondisi       string    `json:"kondisi"`         // Ringan, Berat, Meninggal
	Gejala        string    `json:"gejala"`          // Gejala yang dialami
	TekananDarah string    `json:"tekanan_darah"`   // Tekanan darah korban
	Nadi          int       `json:"nadi"`            // Denyut nadi
	SuhuTubuh     float64   `json:"suhu_tubuh"`      // Suhu tubuh dalam Celcius
	Kesadaran    string    `json:"kesadaran"`       // Sadar, Tidak Sadar, dl
	Deskripsi     string    `json:"deskripsi"`
	WaktuPemeriksaan time.Time `json:"waktu_pemeriksaan"` // Waktu pemeriksaan kondisi
	DilaporkanOleh uint     `json:"dilaporkan_oleh"`  // FK ke petugas atau user
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

