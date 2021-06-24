package db

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/database/redshift"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Migrate(log *zap.Logger, db *gorm.DB, dialect, migrationsPath string) {
	var driver database.Driver
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal("db.Ping", zap.Error(err))
	}

	switch dialect {
	case "postgres":
		driver, err = postgres.WithInstance(sqlDb, &postgres.Config{})
	case "redshift":
		driver, err = redshift.WithInstance(sqlDb, &redshift.Config{})
	default:
		log.Fatal("unexpected dialect " + dialect)
	}

	if err != nil {
		log.Fatal("mysql.WithInstance", zap.Error(err))
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsPath,
		dialect,
		driver,
	)
	if err != nil {
		log.Fatal("migrate.NewWithDatabaseInstance", zap.Error(err))
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("m.Up", zap.Error(err))
	}
}
