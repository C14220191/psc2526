package models

import "time"

type DetailPrivilegeRole struct {
	ID          uint      `json:"id"`
	RoleID      uint      `json:"role_id"`
	PrivilegeID uint      `json:"privilege_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
