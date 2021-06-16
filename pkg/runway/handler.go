package runway

import (
	"dex-trades-parser/pkg/handler"
	"github.com/spf13/viper"
)

func (r *Runway) RootHandler() *handler.Root {
	if !r.flag.IsApp() {
		r.log.Fatal("runway: required app flags")
	}

	return handler.NewRoot(viper.GetString("app-name"))
}
