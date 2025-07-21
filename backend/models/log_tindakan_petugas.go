package models

import "time"

type LogTindakanPetugas struct {
	ID         uint      `json:"id"`
	KasusID    uint      `json:"kasus_id"`
	PetugasID  uint      `json:"petugas_id"`
	Rincian    string    `json:"rincian"`
	WaktuTindakan      time.Time `json:"waktu"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

