package flag

import "github.com/spf13/pflag"

func (f *Flag) IsDB() bool {
	return f.db
}

func (f *Flag) DB() *Flag {
	f.db = true

	pflag.String("db-dsn", "", "Database data source name")
	pflag.String("db-dialect", "postgres", "GORM database dialect")
	pflag.Int("db-max-open-conns", 80, "GORM Max database connections")
	pflag.Bool("db-debug", false, "Debug database")
	pflag.String("migrations-path", "migrations.d", "Path to migrations directory")

	return f
}
