package main

import (
	"context"
	"dex-trades-parser/internal/app"
	"dex-trades-parser/pkg/flag"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

// @title Dexe Investing Api
// @version 1.0
// @description Dexe Investing Api
// @BasePath /

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
