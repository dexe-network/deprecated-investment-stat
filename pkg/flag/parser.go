package flag

import "github.com/spf13/pflag"

func (f *Flag) IsParser() bool {
	return f.parser
}

func (f *Flag) Parser() *Flag {
	f.parser = true

	pflag.Int64("max-parallel-blocks", 200, "Max processing blocks in parallel")
	pflag.String("dex-factory-address", "", "Dex Factory address")
	pflag.String("network", "bsc", "Blockchain network")

	return f
}
