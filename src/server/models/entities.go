package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model

	UserID       uint          `json:"user_id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Transactions []Transaction `json:"transactions,omitempty" gorm:"serializer:json"`
}

type Category struct {
	gorm.Model

	UserID       uint          `json:"user_id,omitempty"`
	Name         string        `json:"name,omitempty"`
	ParentID     *uint         `json:"parent_id,omitempty" gorm:"default:null"`
	Transactions []Transaction `json:"transactions,omitempty" gorm:"serializer:json"`
	Rules        []Rule        `json:"rules,omitempty" gorm:"serializer:json"`
}

type Merchant struct {
	gorm.Model

	UserID       uint          `json:"user_id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Transactions []Transaction `json:"transactions,omitempty" gorm:"serializer:json"`
	Rules        []Rule        `json:"rules,omitempty" gorm:"serializer:json"`
}

type Transaction struct {
	gorm.Model

	UserID      uint      `json:"user_id,omitempty"`
	AccountID   *uint     `json:"account_id,omitempty" gorm:"default:null"`
	Date        time.Time `json:"date,omitempty"`
	Description string    `json:"description,omitempty"`
	Amount      float64   `json:"amount,omitempty"`
	CategoryID  *uint     `json:"category_id,omitempty" gorm:"default:null"`
	MerchantID  *uint     `json:"merchant_id,omitempty" gorm:"default:null"`
	Notes       *string   `json:"notes,omitempty"`
}

type Rule struct {
	gorm.Model

	UserID     uint   `json:"user_id,omitempty"`
	Enabled    bool   `json:"enabled,omitempty" gorm:"default:false"`
	Pattern    string `json:"pattern,omitempty"`
	CategoryID *uint  `json:"category_id,omitempty" gorm:"default:null"`
	MerchantID *uint  `json:"merchant_id,omitempty" gorm:"default:null"`
}
