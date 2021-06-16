package runway

import (
	"dex-trades-parser/pkg/flag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type Runway struct {
	log  *zap.Logger
	flag *flag.Flag
}

const zapTimeLayout = "15:04:05"

func zapTimeEnc(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(zapTimeLayout))
}

func zapProdTimeEnc(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02T15:04:05.000000000Z"))
}

func NewRunway(fl *flag.Flag) *Runway {
	var config zap.Config
	if viper.GetString("app-env") == flag.AppEnvProd {
		config = zap.NewProductionConfig()
		config.EncoderConfig.EncodeTime = zapProdTimeEnc
	} else {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		config.EncoderConfig.EncodeTime = zapTimeEnc
	}
	config.Level = zap.NewAtomicLevelAt(zapcore.Level(viper.GetInt("zap-level")))
	config.OutputPaths = []string{"stderr"}
	config.ErrorOutputPaths = []string{"stderr"}

	log, err := config.Build()
	if err != nil {
		panic(err)
	}
	return &Runway{
		flag: fl,
		log:  log,
	}
}
