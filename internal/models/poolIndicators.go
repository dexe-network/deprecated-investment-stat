package models

import "time"

type PoolIndicators struct {
	Id          uint      `json:"id" gorm:"primaryKey;column:id"`
	PoolAdr     string    `json:"poolAdr"   gorm:"type:character varying(255);column:poolAdr;not null"`
	TotalCap    string    `json:"totalCap"  gorm:"type:numeric;column:totalCap;not null"`
	TotalSupply string    `json:"totalSupply"  gorm:"type:numeric;column:totalSupply;not null"`
	Date        time.Time `json:"date"         gorm:"type:timestamp with time zone;column:date;not null"`
	BlockNumber int64     `json:"blockNumber"  gorm:"type:bigint;column:blockNumber;not null"`                    //blockNumber bigint,
	Tx          string    `json:"tx"           gorm:"type:character varying(255);column:tx;not null;uniqueIndex"` //tx character varying(255),
	CreatedAt   time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}
