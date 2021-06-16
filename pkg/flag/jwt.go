package flag

import (
	"time"

	"github.com/spf13/pflag"
)

func (f *Flag) IsJWT() bool {
	return f.jwt
}

func (f *Flag) JWT() *Flag {
	f.jwt = true

	pflag.Bool("secure-protocol", true, "Indicator to use secure cookies")
	pflag.String("jwt-secret", "dex-trades-parser", "JWT secret")
	pflag.Duration("jwt-access-token-expiration", time.Minute*30, "JWT AccessToken expiration time")
	pflag.Duration("jwt-refresh-token-expiration", time.Hour*24*60, "JWT RefreshToken expiration time")

	return f
}
