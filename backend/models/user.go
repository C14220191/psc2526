package models

// import "time"

type User struct {
	ID           int     `json:"id"`                   // Primary key
	Id_User      string  `json:"id_user"`              // Unique identifier for the user
	Username     string  `json:"username"`             // Unique
	Password     string  `json:"password"`             // Tersimpan terenkripsi
	NamaLengkap  string  `json:"nama_lengkap"`         //
	NIK          string  `json:"nik"`                  // Nomor Induk Kependudukan
	NoTelepon    string  `json:"no_telepon"`           //
	Alamat       string  `json:"alamat"`               //
	Kota         string  `json:"kota"`                 //
	JenisKelamin string  `json:"jenis_kelamin"`        // L/P
	TanggalLahir string  `json:"tanggal_lahir"`        //
	Status       string  `json:"status"`               // aktif, nonaktif, dsb
	OTPStatus    bool    `json:"otp_status"`           // true/false
	CreatedAt    string  `json:"created_at"`           //
	UpdatedAt    string  `json:"updated_at"`           //
	DeletedAt    *string `json:"deleted_at,omitempty"` // soft delete, NULLABLE
}
