package models

import "time"

type Tindakan struct {
	ID         uint      `json:"id"`
	KasusID    uint      `json:"kasus_id"`
	PetugasID  uint      `json:"petugas_id"`
	Jenis      string    `json:"jenis"`     // Pertolongan pertama, evakuasi, rujukan
	Rincian    string    `json:"rincian"`
	Waktu      time.Time `json:"waktu"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
