package models

import (
    "database/sql"
    "time"
)

type Role struct {
    ID        uint         `json:"id_role"`
    NamaRole  string       `json:"nama_role"`
    CreatedAt time.Time    `json:"created_at"`
    UpdatedAt time.Time    `json:"updated_at"`
    DeletedAt sql.NullTime `json:"deleted_at,omitempty"`
}