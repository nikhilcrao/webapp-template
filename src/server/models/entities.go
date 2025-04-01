package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model

	AppUserID    uint          `json:"user_id,omitempty"`
	AppUser      *AppUser      `json:"app_user,omitempty" gorm:"foreignKey:AppUserID"`
	Name         string        `json:"name,omitempty"`
	Transactions []Transaction `json:"transactions,omitempty" gorm:"serializer:json"`
}

type Category struct {
	gorm.Model

	AppUserID uint `json:"user_id,omitempty"`
	//	AppUser         AppUser          `json:"app_user,omitempty"`
	Name     string `json:"name,omitempty"`
	ParentID *uint  `json:"parent_id,omitempty"`
	//	Parent       *Category     `json:"parent,omitempty"`
	Transactions []Transaction `json:"transactions,omitempty" gorm:"serializer:json"`
	Rules        []Rule        `json:"rules,omitempty" gorm:"serializer:json"`
}

type Merchant struct {
	gorm.Model

	AppUserID uint `json:"user_id,omitempty"`
	//	AppUser         AppUser          `json:"app_user,omitempty"`
	Name         string        `json:"name,omitempty"`
	Transactions []Transaction `json:"transactions,omitempty" gorm:"serializer:json"`
	Rules        []Rule        `json:"rules,omitempty" gorm:"serializer:json"`
}

type Transaction struct {
	gorm.Model

	AppUserID uint `json:"user_id,omitempty"`
	//	AppUser        AppUser      `json:"app_user,omitempty"`
	AccountID uint `json:"account_id,omitempty"`
	//	Account     Account   `json:"account,omitempty"`
	Date        time.Time `json:"date,omitempty"`
	Description string    `json:"description,omitempty"`
	Amount      float64   `json:"amount,omitempty"`
	CategoryID  uint      `json:"category_id,omitempty"`
	//	Category    Category  `json:"category,omitempty"`
	MerchantID uint `json:"merchant_id,omitempty"`
	//	Merchant    Merchant  `json:"merchant,omitempty"`
	Notes string `json:"notes,omitempty"`
}

type Rule struct {
	gorm.Model

	AppUserID uint `json:"user_id,omitempty"`
	//	AppUser       AppUser   `json:"app_user,omitempty"`
	Enabled    bool   `json:"enabled,omitempty"`
	Pattern    string `json:"pattern,omitempty"`
	CategoryID uint   `json:"category_id,omitempty"`
	//	Category   Category `json:"category,omitempty"`
	MerchantID uint `json:"merchant_id,omitempty"`
	// Merchant   Merchant `json:"merchant,omitempty"`
}
