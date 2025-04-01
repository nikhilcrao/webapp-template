package models

import (
	"gorm.io/gorm"
)

type AppUser struct {
	gorm.Model

	Name         string `json:"name"`
	Email        string `json:"email" binding:"email"`
	PasswordHash string `json:"-"` // Never expose password hash
	GoogleID     string `json:"google_id,omitempty"`

	Accounts     []Account     `json:"accounts,omitempty"`
	Categories   []Category    `json:"categories,omitempty"`
	Merchants    []Merchant    `json:"merchants,omitempty"`
	Transactions []Transaction `json:"transactions,omitempty"`
	Rules        []Rule        `json:"rules,omitempty"`
}
