package models

import (
	"time"
)

type PoolTransfer struct {
	Id          uint            `gorm:"primaryKey;column:id"`
	Wallet      string          `json:"wallet"      gorm:"type:character varying(255);column:wallet;not null"`
	PoolAdr     string          `json:"poolAdr"   gorm:"type:character varying(255);column:poolAdr;not null"`
	Amount      string  `json:"amount"   gorm:"type:numeric;column:amount;not null"`
	Type        string `json:"type" gorm:"type:pooltransfertype;column:type;not null"` // enum('deposit', 'withdraw')
	Date        time.Time       `json:"date"         gorm:"type:timestamp with time zone;column:date;not null"`
	BlockNumber int64           `json:"blockNumber"  gorm:"type:bigint;column:blockNumber;not null"`        //blockNumber bigint,
	Tx          string          `json:"tx"           gorm:"type:character varying(255);column:tx;not null;uniqueIndex"` //tx character varying(255),
	CreatedAt   time.Time       `gorm:"column:createdAt"`
	UpdatedAt   time.Time       `gorm:"column:updatedAt"`
}
