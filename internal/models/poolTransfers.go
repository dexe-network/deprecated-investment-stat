package models

import (
	"github.com/jackc/pgtype"
	"time"
)

type PoolTransfer struct {
	ID          uint            `gorm:"primaryKey"`
	Wallet      string          `json:"wallet"      gorm:"type:character varying(255);column:wallet;not null"`
	PoolAdr     string          `json:"poolAdr"   gorm:"type:character varying(255);column:poolAdr;not null"`
	Amount      pgtype.Numeric  `json:"amount"   gorm:"type:numeric;column:amount;not null"`
	Type        pgtype.EnumType `json:"type" gorm:"type:pooltransfertype;column:type;not null"` // enum('deposit', 'withdraw')
	Date        time.Time       `json:"date"         gorm:"type:timestamp with time zone;column:date;not null"`
	BlockNumber int64           `json:"blockNumber"  gorm:"type:bigint;column:blockNumber;not null"`        //blockNumber bigint,
	Tx          string          `json:"tx"           gorm:"type:character varying(255);column:tx;not null"` //tx character varying(255),
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
