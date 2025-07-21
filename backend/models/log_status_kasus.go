package models

import "time"

type LogStatusKasus struct {
    ID            uint      `json:"id_log_status_kasus"`
    IDKasus       uint      `json:"id_kasus"`
    Deskripsi     string    `json:"deskripsi"`
    WaktuTindakan time.Time `json:"waktu_tindakan"`
}