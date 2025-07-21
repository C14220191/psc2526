package models

import "time"

type OTPStatus struct {
	ID         uint      `json:"id"`
	UserID     uint      `json:"user_id"`
	KodeOTP    string    `json:"kode_otp"`
	Status     string    `json:"status"`       // pending, verified, expired
	TerkirimAt time.Time `json:"terkirim_at"`
	ExpiredAt  time.Time `json:"expired_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
