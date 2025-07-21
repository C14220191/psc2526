package models

import "time"

type LaporanObat struct {
	ID         uint      `json:"id"`
	KasusID    uint      `json:"kasus_id"`
	NamaObat   string    `json:"nama_obat"`
	Dosis      string    `json:"dosis"`
	Jumlah     int       `json:"jumlah"`
	CaraPemberian string    `json:"cara_pemberian"` // Oral, Injeksi, dll
	WaktuPemberian time.Time `json:"waktu_pemberian"`
	Keterangan string    `json:"keterangan"` // Catatan tambaha
	PetugasID  uint      `json:"petugas_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

