package flag

import (
	"github.com/spf13/pflag"
)

const (
	AppEnvProd = "prod"
	AppEnvDev  = "dev"
)

func (f *Flag) IsApp() bool {
	return f.app
}

func (f *Flag) App() *Flag {
	f.app = true
	pflag.String("app-env", AppEnvProd, "Application environment")
	pflag.String("app-addr", "0.0.0.0:3000", "Application address")
	return f
}
