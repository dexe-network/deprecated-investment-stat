package runway

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"

	"github.com/spf13/viper"
)

func (r *Runway) ETH() (client *ethclient.Client) {
	if !r.flag.IsEth() {
		r.log.Fatal("runway: required eth flags")
	}


	client, err := ethclient.Dial(viper.GetString("eth-node"))
	if err != nil {
		r.log.Fatal("ethclient.Dial",
			zap.Error(err),
			zap.String("eth-node", viper.GetString("eth-node")),
		)
	}

	return
}
