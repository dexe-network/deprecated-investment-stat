package flag

import (
	"strings"

	_ "github.com/joho/godotenv/autoload"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Flag struct {
	app    bool
	db     bool
	jwt    bool
	eth    bool
	debug  bool
	parser bool
	cors   bool
}

func New() *Flag {
	return &Flag{}
}

func (f *Flag) Parse() (err error) {
	pflag.Parse()

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	return viper.BindPFlags(pflag.CommandLine)
}
