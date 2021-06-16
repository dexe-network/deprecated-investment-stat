package main

import (
	"context"
	"dex-trades-parser/internal/app"
	"dex-trades-parser/pkg/flag"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fl := flag.New()
	fl.App().
		CORS().
		DB().
		ETH().
		Debug().
		Parser()

	if err := fl.Parse(); err != nil {
		panic(err)
	}

	app.Run(ctx, cancel, fl)
}
