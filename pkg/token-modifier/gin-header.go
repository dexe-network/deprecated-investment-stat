package token_modifier

import (
	"dex-trades-parser/pkg/jwtoken"
	"github.com/gin-gonic/gin"
)

const (
	GinHeaderAccessToken = "Authorization"
)

type GinHeaders struct {
	c *gin.Context
}

func NewGinHeaders(c *gin.Context) GinHeaders {
	return GinHeaders{
		c: c,
	}
}

func (m GinHeaders) Get() (tokens jwtoken.Tokens) {
	return jwtoken.Tokens{
		AccessToken: m.c.GetHeader(GinHeaderAccessToken),
	}
}

func (m GinHeaders) Set(tokens jwtoken.Tokens) {
	m.c.Header(GinHeaderAccessToken, tokens.AccessToken)
}

func (m GinHeaders) Remove() {}
