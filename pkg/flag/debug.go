package flag

import (
	"github.com/spf13/pflag"
	"go.uber.org/zap"
)

func (f *Flag) IsDebug() bool {
	return f.debug
}

func (f *Flag) Debug() *Flag {
	f.debug = true

	pflag.Int8("zap-level", int8(zap.DebugLevel), "debug -1, info 0, warn 1, error 2, dpanic 3, panic 4, fatal 5")

	return f
}
