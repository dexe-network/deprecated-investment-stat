package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type EthTrade struct {
	TokenA		string				`json:"tokenA"       gorm:"type:character varying(255);column:tokenA"`		//tokenA character varying(255),
	TokenB		string				`json:"tokenB"       gorm:"type:character varying(255);column:tokenB"`		//tokenB character varying(255),
	Date		time.Time			`json:"date"         gorm:"type:timestamp with time zone;column:date"`		//date timestamp with time zone,
	BlockNumber	int64				`json:"blockNumber"  gorm:"type:bigint;column:blockNumber"`					//blockNumber bigint,
	Tx			string				`json:"tx"           gorm:"type:character varying(255);column:tx"`			//tx character varying(255),
	Protocol	string				`json:"protocol"     gorm:"type:character varying(255);column:protocol"`	//protocol character varying(255),
	PriceIn		decimal.Decimal		`json:"priceIn"      gorm:"type:numeric;column:priceIn"`					//priceIn numeric,
	PriceOut	decimal.Decimal		`json:"priceOut"     gorm:"type:numeric;column:priceOut"`					//priceOut numeric,
	AmountOut	string				`json:"amountOut"    gorm:"type:character varying(255);column:amountOut"`	//amountOut character varying(255),
	AmountIn	string				`json:"amountIn"     gorm:"type:character varying(255);column:amountIn"`	//amountIn character varying(255),
	Wallet		string				`json:"wallet"       gorm:"type:character varying(255);column:wallet"`		//wallet character varying(255),
	Value		string				`json:"value"        gorm:"type:character varying(255);column:value"`		//value character varying(255)
}
