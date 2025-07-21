package models

import "database/sql"

type Petugas struct {
    ID                uint           `json:"id_petugas"`
    IDFaskes          uint           `json:"id_faskes"`
    IDTindakanPetugas sql.NullInt64  `json:"id_tindakan_petugas"`
    Nama              string         `json:"nama"`
    JenisKelamin      string         `json:"jenis_kelamin"`
    Profesi           string         `json:"profesi"`
    Username          string         `json:"username"`
    Password          string         `json:"-"`
    NoHP              string         `json:"no_hp"`
    Keahlian          sql.NullString `json:"keahlian"`
}