package runway

import (
	"dex-trades-parser/pkg/parser"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func (r *Runway) Parser() (p *parser.Parser) {
	if !r.flag.IsParser() {
		r.log.Fatal("runway: required parser flags")
	}

	p, err := parser.NewParser(r.log, r.ETH(), viper.GetString("dex-router-address"), viper.GetString("dex-router-abi"), viper.GetString("dex-protocol"))
	if err != nil {
		r.log.Fatal("parser.NewParser",
			zap.Error(err),
			zap.String("dex-router-abi", viper.GetString("dex-router-abi")),
			zap.String("dex-router-address", viper.GetString("dex-router-address")),
		)
	}

	return
}
