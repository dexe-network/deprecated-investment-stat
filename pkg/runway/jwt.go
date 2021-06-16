package runway

import (
	"time"

	"github.com/spf13/viper"

	"dex-trades-parser/pkg/jwtoken"
	token_modifier "dex-trades-parser/pkg/token-modifier"
)

func (r *Runway) JWT() *jwtoken.JWT {
	if !r.flag.IsJWT() {
		r.log.Fatal("runway: required jwt flags")
	}

	return jwtoken.NewJWT(r.log, jwtoken.JWTConfig{
		Secret:                 []byte(viper.GetString("jwt-secret")),
		AccessTokenExpiration:  viper.GetDuration("jwt-access-token-expiration"),
		RefreshTokenExpiration: viper.GetDuration("jwt-refresh-token-expiration"),
	})
}

func (r *Runway) TokenModifier() token_modifier.Gin {
	if !r.flag.IsJWT() {
		r.log.Fatal("runway: required jwt flags")
	}

	return token_modifier.NewGin(time.Hour*24, viper.GetBool("secure-protocol"))
}
