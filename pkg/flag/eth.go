package flag

import "github.com/spf13/pflag"

func (f *Flag) IsEth() bool {
	return f.eth
}

func (f *Flag) ETH() *Flag {
	f.eth = true

	pflag.String("eth-node", "ws://127.0.0.1:8546", "Eth node URL")

	return f
}
