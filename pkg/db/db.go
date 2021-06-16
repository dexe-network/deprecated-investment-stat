package db

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/database/redshift"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func Migrate(log *zap.Logger, db *gorm.DB, dialect, migrationsPath string) {
	var driver database.Driver
	err := db.DB().Ping()
	if err != nil {
		log.Fatal("db.Ping", zap.Error(err))
	}

	switch dialect {
	case "postgres":
		driver, err = postgres.WithInstance(db.DB(), &postgres.Config{})
	case "redshift":
		driver, err = redshift.WithInstance(db.DB(), &redshift.Config{})
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
