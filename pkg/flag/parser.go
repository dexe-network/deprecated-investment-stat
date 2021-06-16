package flag

import "github.com/spf13/pflag"

func (f *Flag) IsParser() bool {
	return f.parser
}

func (f *Flag) Parser() *Flag {
	f.parser = true

	pflag.String("dex-router-address", "", "Dex Router address")
	pflag.String("dex-router-abi", "", "Dex Router ABI")
	pflag.String("dex-protocol", "uniswapV2", "Dex Protocol")
	pflag.Int64("max-parallel-blocks", 200, "Max processing blocks in parallel")

	return f
}
