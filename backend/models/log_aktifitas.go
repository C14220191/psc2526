package models

import "time"

type LogAktivitas struct {
    ID        uint      `json:"id_log"`
    IDAdmin   uint      `json:"id_admin"`
    IDPetugas uint      `json:"id_petugas"`
    Aksi      string    `json:"aksi"`
    Deskripsi string    `json:"deskripsi"`
    Waktu     time.Time `json:"waktu"`
}