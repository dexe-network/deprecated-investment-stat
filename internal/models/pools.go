package models

import (
	"time"
)

type Pool struct {
	Id                    uint      `json:"id" gorm:"primaryKey;column:id"`
	CreatorAdr            string    `json:"creatorAdr"      gorm:"type:character varying(255);column:creatorAdr;not null"`
	BasicTokenAdr         string    `json:"basicTokenAdr"   gorm:"type:character varying(255);column:basicTokenAdr;not null"`
	TotalSupply           string    `json:"totalSupply"  gorm:"type:numeric;column:totalSupply;not null"`
	TraderCommissionNum   uint16    `json:"traderCommissionNum" gorm:"type:integer;column:traderCommissionNum;not null"`
	TraderCommissionDen   uint16    `json:"traderCommissionDen" gorm:"type:integer;column:traderCommissionDen;not null"`
	InvestorCommissionNum uint16    `json:"investorCommissionNum" gorm:"type:integer;column:investorCommissionNum;not null"`
	InvestorCommissionDen uint16    `json:"investorCommissionDen" gorm:"type:integer;column:investorCommissionDen;not null"`
	DexeCommissionNum     uint16    `json:"dexeCommissionNum" gorm:"type:integer;column:dexeCommissionNum;not null"`
	DexeCommissionDen     uint16    `json:"dexeCommissionDen" gorm:"type:integer;column:dexeCommissionDen;not null"`
	IsActualOn            bool      `json:"isActualOn"      gorm:"type:boolean;column:isActualOn;not null"`
	InvestorRestricted    bool      `json:"investorRestricted"      gorm:"type:bool;column:investorRestricted;not null"`
	Name                  string    `json:"name"      gorm:"type:character varying(255);column:name;not null"`
	Symbol                string    `json:"symbol"      gorm:"type:character varying(255);column:symbol;not null"`
	PoolAdr               string    `json:"poolAdr"   gorm:"type:character varying(255);column:poolAdr;not null;uniqueIndex"`
	Date                  time.Time `json:"date"         gorm:"type:timestamp with time zone;column:date;not null"`
	BlockNumber           int64     `json:"blockNumber"  gorm:"type:bigint;column:blockNumber;not null"`                    //blockNumber bigint,
	Tx                    string    `json:"tx"           gorm:"type:character varying(255);column:tx;not null;uniqueIndex"` //tx character varying(255),
	CreatedAt             time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt             time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}
