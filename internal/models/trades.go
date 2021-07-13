package models

import (
	"github.com/jackc/pgtype"
	"time"
)

type Trade struct {
	Id           uint             `gorm:"primaryKey;column:id"`
	TraderPool   string           `json:"traderPool"      gorm:"type:character varying(255);column:traderPool;not null"`
	AmountIn     string   `json:"amountIn"   gorm:"type:numeric;column:amountIn;not null"`
	AmountOutMin string   `json:"amountOutMin"  gorm:"type:numeric;column:amountOutMin;not null"`
	Path         pgtype.TextArray `json:"path" gorm:"type:text[];column:path;not null"`
	Deadline     string   `json:"deadline" gorm:"type:numeric;column:deadline;not null"`
	Date         time.Time        `json:"date"         gorm:"type:timestamp with time zone;column:date;not null"`
	BlockNumber  int64            `json:"blockNumber"  gorm:"type:bigint;column:blockNumber;not null"`        //blockNumber bigint,
	Tx           string           `json:"tx"           gorm:"type:character varying(255);column:tx;not null;uniqueIndex"` //tx character varying(255),
	CreatedAt    time.Time        `gorm:"column:createdAt"`
	UpdatedAt    time.Time        `gorm:"column:updatedAt"`
}
