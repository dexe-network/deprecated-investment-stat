package models

import (
	"time"
)

type TradeItem struct {
	Id                uint         `json:"id" gorm:"primaryKey;column:id"`
	PoolAddress       string       `json:"poolAddress"      gorm:"type:character varying(255);column:poolAddress;not null"`
	BaseTokenAddress  string       `json:"baseTokenAddress"      gorm:"type:character varying(255);column:baseTokenAddress;not null"`
	TradeTokenAddress string       `json:"tradeTokenAddress"      gorm:"type:character varying(255);column:tradeTokenAddress;not null"`
	Balance           string       `json:"balance"      gorm:"type:numeric;column:balance;not null"`
	TradeStatus       string       `json:"tradeStatus"      gorm:"type:tradestatus;column:tradeStatus;not null"`
	TradeEvents       []TradeEvent `json:"tradeEvents" gorm:"foreignKey:trade_item_id"`

	OpenDate    time.Time `json:"openDate"         gorm:"type:timestamp with time zone;column:openDate;not null"`
	CloseDate   time.Time `json:"closeDate"         gorm:"type:timestamp with time zone;column:closeDate;not null"`
	Date        time.Time `json:"date"         gorm:"type:timestamp with time zone;column:date;not null"`
	BlockNumber int64     `json:"blockNumber"  gorm:"type:bigint;column:blockNumber;not null"` //blockNumber bigint,
	CreatedAt   time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}
