package models

import "time"

type User struct {
	ID           uint      `json:"id"`                      // Primary key
	Id_User      string    `json:"id_user"`                 // Unique identifier for the user
	Username     string    `json:"username"`                // Unique
	Password     string    `json:"password"`                // Tersimpan terenkripsi
	NamaLengkap  string    `json:"nama_lengkap"`            //
	NIK          string    `json:"nik"`                     // Nomor Induk Kependudukan
	NoTelepon    string    `json:"no_telepon"`              //
	Alamat       string    `json:"alamat"`                  //
	Kota         string    `json:"kota"`                    //
	JenisKelamin string    `json:"jenis_kelamin"`           // L/P
	TanggalLahir time.Time `json:"tanggal_lahir"`           //
	Status       string    `json:"status"`                  // aktif, nonaktif, dsb
	OTPStatus    bool      `json:"otp_status"`              // true/false
	CreatedAt    time.Time `json:"created_at"`              //
	UpdatedAt    time.Time `json:"updated_at"`              //
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`   // soft delete, NULLABLE
}
