package models

import (
    "database/sql"
    "time"
)

type Berita struct {
    ID                   uint           `json:"id_berita"`
    IDAdminPembuatBerita uint           `json:"id_admin_pembuat_berita"`
    Judul                string         `json:"judul"`
    Isi                  string         `json:"isi"`
    Thumbnail            sql.NullString `json:"thumbnail"`
    Tanggal              sql.NullTime   `json:"tanggal_date"`
    CreatedAt            time.Time      `json:"created_at"`
    UpdatedAt            time.Time      `json:"updated_at"`
    DeletedAt            sql.NullTime   `json:"deleted_at,omitempty"`
}