package models

import (
	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Trade struct {
	TraderPool   string           `json:"traderPool"      gorm:"type:character varying(255);column:traderPool;not null"`
	AmountIn     pgtype.Numeric   `json:"amountIn"   gorm:"type:numeric;column:amountIn;not null"`
	AmountOutMin pgtype.Numeric   `json:"amountOutMin"  gorm:"type:numeric;column:amountOutMin;not null"`
	Path         pgtype.TextArray `json:"path" gorm:"type:text[];column:path;not null"`
	Deadline     pgtype.Numeric   `json:"deadline" gorm:"type:numeric;column:deadline;not null"`
	BlockNumber  int64            `json:"blockNumber"  gorm:"type:bigint;column:blockNumber;not null"`        //blockNumber bigint,
	Tx           string           `json:"tx"           gorm:"type:character varying(255);column:tx;not null"` //tx character varying(255),
	gorm.Model
}
