package runway

import (
	"dex-trades-parser/migrations"
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (r *Runway) DB() (db *gorm.DB) {
	if !r.flag.IsDB() {
		r.log.Fatal("runway: required db flags")
	}

	db, dbErr := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			viper.GetString("db-host"),
			viper.GetString("db-port"),
			viper.GetString("db-user"),
			viper.GetString("db-password"),
			viper.GetString("db-dbname")),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		// Need move to env
		Logger: logger.Default.LogMode(1),
	})

	if dbErr != nil {
		r.log.Fatal("gorm.Open",
			zap.Error(dbErr),
			zap.String("db-dsn", viper.GetString("db-dsn")),
		)
	}

	dbConnectionPool, cpErr := db.DB()
	if cpErr != nil {
		r.log.Fatal("gorm.Open",
			zap.Error(cpErr),
			zap.String("db-dsn", viper.GetString("db-dsn")),
		)
	}
	dbConnectionPool.SetMaxOpenConns(viper.GetInt("db-max-open-conns"))

	// Migrations
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		migrations.InitTables_201608301400,
	})

	if dbErr = m.Migrate(); dbErr != nil {
		r.log.Fatal("Could not migrate", zap.Error(dbErr))
	}
	fmt.Println("Migration did run successfully")

	return
}
