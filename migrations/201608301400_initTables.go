package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var InitTables_201608301400 = &gormigrate.Migration{
	ID: "201608301400",
	Migrate: func(tx *gorm.DB) error {
		// it's a good pratice to copy the struct inside the function,
		// so side effects are prevented if the original struct changes during the time
		type Pool struct {
			CreatorAdr            string `json:"CreatorAdr"      gorm:"type:character varying(255);column:CreatorAdr;not null"`
			BasicTokenAdr         string `json:"BasicTokenAdr"   gorm:"type:character varying(255);column:BasicTokenAdr;not null"`
			TotalSupply           string `json:"TotalSupply"  gorm:"type:character varying(255);column:TotalSupply;not null"`
			TraderCommissionNum   int16  `json:"TraderCommissionNum" gorm:"type:integer;column:TraderCommissionNum;not null"`
			TraderCommissionDen   int16  `json:"TraderCommissionDen" gorm:"type:integer;column:TraderCommissionDen;not null"`
			InvestorCommissionNum int16  `json:"InvestorCommissionNum" gorm:"type:integer;column:InvestorCommissionNum;not null"`
			InvestorCommissionDen int16  `json:"InvestorCommissionDen" gorm:"type:integer;column:InvestorCommissionDen;not null"`
			DexeCommissionNum     int16  `json:"DexeCommissionNum" gorm:"type:integer;column:DexeCommissionNum;not null"`
			DexeCommissionDen     int16  `json:"DexeCommissionDen" gorm:"type:integer;column:DexeCommissionDen;not null"`
			IsActualOn            bool   `json:"IsActualOn"      gorm:"type:boolean;column:IsActualOn;not null"`
			InvestorRestricted    bool   `json:"InvestorRestricted"      gorm:"type:bool;column:InvestorRestricted;not null"`
			Name                  string `json:"Name"      gorm:"type:character varying(255);column:Name;not null"`
			Symbol                string `json:"Symbol"      gorm:"type:character varying(255);column:Symbol;not null"`
			PoolAdr               string `json:"PoolAdr"   gorm:"type:character varying(255);column:PoolAdr;not null"`
			BlockNumber           int64  `json:"blockNumber"  gorm:"type:bigint;column:blockNumber;not null"`        //blockNumber bigint,
			Tx                    string `json:"tx"           gorm:"type:character varying(255);column:tx;not null"` //tx character varying(255),
			gorm.Model
		}
		return tx.AutoMigrate(&Pool{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("pools")
	},
}
