package models

import (
	"time"
)

type GlobalTokenWhitelist struct {
	ID          uint      `gorm:"primaryKey"`
	Address     string    `json:"address"      gorm:"type:character varying(255);column:address;not null"`
	Date        time.Time `json:"date"         gorm:"type:timestamp with time zone;column:date;not null"`
	BlockNumber int64     `json:"blockNumber"  gorm:"type:bigint;column:blockNumber;not null"`        //blockNumber bigint,
	Tx          string    `json:"tx"           gorm:"type:character varying(255);column:tx;not null"` //tx character varying(255),
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
