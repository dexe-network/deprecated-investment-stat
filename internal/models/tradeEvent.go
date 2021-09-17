package models

import (
	"time"
)

type TradeEvent struct {
	Id                   uint   `json:"id" gorm:"primaryKey;column:id"`
	TradeType            string `json:"tradeType"      gorm:"type:tradetype;column:tradeType;not null"`
	SpentTokenAddress    string `json:"spentTokenAddress"      gorm:"type:character varying(255);column:spentTokenAddress;not null"`
	ReceivedTokenAddress string `json:"receivedTokenAddress"      gorm:"type:character varying(255);column:receivedTokenAddress;not null"`
	Price                string `json:"price"  gorm:"type:numeric;column:price;not null"`
	FromAmount           string `json:"fromAmount"      gorm:"type:numeric;column:fromAmount;not null"`
	ToAmount             string `json:"toAmount"      gorm:"type:numeric;column:toAmount;not null"`
	TradeItemID          uint
	Date                 time.Time `json:"date"         gorm:"type:timestamp with time zone;column:date;not null"`         // enum('buy', 'sell')
	BlockNumber          int64     `json:"blockNumber"  gorm:"type:bigint;column:blockNumber;not null"`                    //blockNumber bigint,
	Tx                   string    `json:"tx"           gorm:"type:character varying(255);column:tx;not null;uniqueIndex"` //tx character varying(255),
	CreatedAt            time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt            time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}
