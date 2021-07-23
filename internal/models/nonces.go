package models

import (
	"time"
)

type Nonce struct {
	Id        uint      `json:"id" gorm:"primaryKey;column:id"`
	Wallet    string    `json:"wallet"      gorm:"type:character varying(255);column:wallet;not null;uniqueIndex"`
	Nonce     int       `json:"nonce"  gorm:"type:bigint;column:nonce;not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}
