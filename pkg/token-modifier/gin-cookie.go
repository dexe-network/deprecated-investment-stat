package token_modifier

import (
	"time"

	"dex-trades-parser/pkg/jwtoken"
	"github.com/gin-gonic/gin"
)

const (
	GinCookieAccessToken = "authorization"
)

type GinCookies struct {
	c      *gin.Context
	maxAge int
	secure bool
}

func NewGinCookies(c *gin.Context, maxAge time.Duration, secure bool) GinCookies {
	return GinCookies{
		c:      c,
		maxAge: int(maxAge / time.Second),
		secure: secure,
	}
}

func (m GinCookies) Get() (tokens jwtoken.Tokens) {
	at, _ := m.c.Cookie(GinCookieAccessToken)

	return jwtoken.Tokens{
		AccessToken: at,
	}
}

func (m GinCookies) Set(tokens jwtoken.Tokens) {
	m.c.SetCookie(GinCookieAccessToken, tokens.AccessToken, m.maxAge, "/", "", m.secure, true)
}

func (m GinCookies) Remove() {
	m.c.SetCookie(GinCookieAccessToken, "", -1, "/", "", m.secure, true)
}
