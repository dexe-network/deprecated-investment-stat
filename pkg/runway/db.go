package runway

import (
	pkgdb "dex-trades-parser/pkg/db"
	"go.uber.org/zap"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func (r *Runway) DB() (db *gorm.DB) {
	if !r.flag.IsDB() {
		r.log.Fatal("runway: required db flags")
	}


	db, err := gorm.Open("postgres", viper.GetString("db-dsn"))
	db.DB().SetMaxOpenConns(viper.GetInt("db-max-open-conns"))


	if err != nil {
		r.log.Fatal("gorm.Open",
			zap.Error(err),
			zap.String("db-dsn", viper.GetString("db-dsn")),
		)
	}

	db.LogMode(viper.GetBool("db-debug"))

	if viper.GetString("migrations-path") != "" {
		pkgdb.Migrate(r.log, db, viper.GetString("db-dialect"), viper.GetString("migrations-path"))
	}

	return
}
