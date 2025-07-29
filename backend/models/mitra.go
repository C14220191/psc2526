// models/mitra.go
package models

import "time"

type Mitra struct {
	ID        uint      `json:"id"`
	Nama      string    `json:"nama" validate:"required"`
	Alamat    string    `json:"alamat" validate:"required"`
	Kontak    string    `json:"kontak" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MitraCreate struct {
	Nama   string `json:"nama" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
	Kontak string `json:"kontak" validate:"required"`
	Email  string `json:"email" validate:"required,email"`
}

type MitraUpdate struct {
	ID     uint   `json:"id" validate:"required"`
	Nama   string `json:"nama" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
	Kontak string `json:"kontak" validate:"required"`
	Email  string `json:"email" validate:"required,email"`
}

type MitraFilter struct {
	PageNumber int    `json:"page_number"`
	PageSize   int    `json:"page_size"`
	Nama       string `json:"nama"`
	Email      string `json:"email"`
	Kontak     string `json:"kontak"`
}

// type Response struct {
// 	StatusCode int         `json:"StatusCode"`
// 	Message    string      `json:"Message"`
// 	Data       interface{} `json:"Data"`
// }
