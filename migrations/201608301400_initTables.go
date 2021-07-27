package migrations

import (
	"dex-trades-parser/internal/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var InitTables_201608301400 = &gormigrate.Migration{
	ID: "201608301400",
	Migrate: func(tx *gorm.DB) error {
		// it's a good pratice to copy the struct inside the function,
		// so side effects are prevented if the original struct changes during the time
		tx.Exec("CREATE TYPE pooltransfertype AS ENUM('deposit', 'withdraw')")
		return tx.AutoMigrate(
			&models.Pool{},
			&models.Trade{},
			&models.PoolTransfer{},
			&models.GlobalTokenWhitelist{},
			&models.Nonce{},
			&models.User{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("pools", "trades")
	},
}
