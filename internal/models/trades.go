package models

import (
	"github.com/jackc/pgtype"
	"time"
)

type Trade struct {
	Id           uint             `json:"id" gorm:"primaryKey;column:id"`
	TraderPool   string           `json:"traderPool"      gorm:"type:character varying(255);column:traderPool;not null"`
	AmountIn     string           `json:"amountIn"   gorm:"type:numeric;column:amountIn;not null"`        // Input Data
	AmountOutMin string           `json:"amountOutMin"  gorm:"type:numeric;column:amountOutMin;not null"` // Input Data
	FromAmt      string           `json:"fromAmt"   gorm:"type:numeric;column:fromAmt;not null"`          // Logs Data - real amount spend
	ToAmt        string           `json:"toAmt"   gorm:"type:numeric;column:toAmt;not null"`              // Logs Data - real amount receive
	Path         pgtype.TextArray `json:"path" gorm:"type:text[];column:path;not null"`
	Deadline     string           `json:"deadline" gorm:"type:numeric;column:deadline;not null"`
	Date         time.Time        `json:"date"         gorm:"type:timestamp with time zone;column:date;not null"`
	Type         string           `json:"type" gorm:"type:tradetype;column:type;not null"`                                // enum('buy', 'sell')
	BlockNumber  int64            `json:"blockNumber"  gorm:"type:bigint;column:blockNumber;not null"`                    //blockNumber bigint,
	Tx           string           `json:"tx"           gorm:"type:character varying(255);column:tx;not null;uniqueIndex"` //tx character varying(255),
	CreatedAt    time.Time        `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt    time.Time        `json:"updatedAt" gorm:"column:updatedAt"`
}
