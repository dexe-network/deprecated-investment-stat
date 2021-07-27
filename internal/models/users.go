package models

import (
	"time"
)

type User struct {
	Id        uint      `json:"id" gorm:"primaryKey;column:id"`
	Wallet    string    `json:"wallet"      gorm:"type:character varying(255);column:wallet;not null;uniqueIndex"`
	Nickname  string    `json:"nickname"           gorm:"type:character varying(255);column:nickname;not null"`
	Avatar    string    `json:"avatar"           gorm:"type:character varying(255);column:avatar;not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}
