package models

import (
    "database/sql"
    
)

type Admin struct {
    ID          uint           `json:"id_admin"`
    Username    string         `json:"username"`
    Password    string         `json:"-"` // Sembunyikan dari JSON
    NamaLengkap string         `json:"nama_lengkap"`
    Email       string         `json:"email"`
    NoTelepon   sql.NullString `json:"no_telepon"`
}