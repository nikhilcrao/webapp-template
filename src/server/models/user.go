package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name         string `json:"name"`
	Email        string `json:"email" binding:"email"`
	PasswordHash string `json:"-"` // Never expose password hash
	GoogleID     string `json:"google_id,omitempty"`
}
