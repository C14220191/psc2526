package models

import (
    "database/sql"
    "time"
)

type DetailRoleAdmin struct {
    ID        uint         `json:"id_detail_role"`
    IDAdmin   uint         `json:"id_admin"`
    IDRole    uint         `json:"id_role"`
    CreatedAt time.Time    `json:"created_at"`
    UpdatedAt time.Time    `json:"updated_at"`
    DeletedAt sql.NullTime `json:"deleted_at,omitempty"`
}