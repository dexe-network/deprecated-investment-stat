package models

import (
	"gorm.io/gorm"
)

type Pool struct {
	CreatorAdr            string `json:"CreatorAdr"      gorm:"type:character varying(255);column:CreatorAdr;not null"`
	BasicTokenAdr         string `json:"BasicTokenAdr"   gorm:"type:character varying(255);column:BasicTokenAdr;not null"`
	TotalSupply           string `json:"TotalSupply"  gorm:"type:character varying(255);column:TotalSupply;not null"`
	TraderCommissionNum   uint16  `json:"TraderCommissionNum" gorm:"type:integer;column:TraderCommissionNum;not null"`
	TraderCommissionDen   uint16  `json:"TraderCommissionDen" gorm:"type:integer;column:TraderCommissionDen;not null"`
	InvestorCommissionNum uint16  `json:"InvestorCommissionNum" gorm:"type:integer;column:InvestorCommissionNum;not null"`
	InvestorCommissionDen uint16  `json:"InvestorCommissionDen" gorm:"type:integer;column:InvestorCommissionDen;not null"`
	DexeCommissionNum     uint16  `json:"DexeCommissionNum" gorm:"type:integer;column:DexeCommissionNum;not null"`
	DexeCommissionDen     uint16  `json:"DexeCommissionDen" gorm:"type:integer;column:DexeCommissionDen;not null"`
	IsActualOn            bool   `json:"IsActualOn"      gorm:"type:boolean;column:IsActualOn;not null"`
	InvestorRestricted    bool   `json:"InvestorRestricted"      gorm:"type:bool;column:InvestorRestricted;not null"`
	Name                  string `json:"Name"      gorm:"type:character varying(255);column:Name;not null"`
	Symbol                string `json:"Symbol"      gorm:"type:character varying(255);column:Symbol;not null"`
	PoolAdr               string `json:"PoolAdr"   gorm:"type:character varying(255);column:PoolAdr;not null"`
	BlockNumber           int64  `json:"blockNumber"  gorm:"type:bigint;column:blockNumber;not null"`        //blockNumber bigint,
	Tx                    string `json:"tx"           gorm:"type:character varying(255);column:tx;not null"` //tx character varying(255),
	gorm.Model
}
