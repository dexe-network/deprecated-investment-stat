package token_modifier

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	maxAge time.Duration
	secure bool
}

func NewGin(maxAge time.Duration, secure bool) Gin {
	return Gin{
		maxAge: maxAge,
		secure: secure,
	}
}

func (g Gin) Getter(c *gin.Context) Getter {
	return NewGinCombineGetter(
		NewGinHeaders(c),
		NewGinCookies(c, g.maxAge, g.secure),
	)
}

func (g Gin) Setter(c *gin.Context) Setter {
	return NewGinCombineSetter(
		NewGinHeaders(c),
		NewGinCookies(c, g.maxAge, g.secure),
	)
}

func (g Gin) Remover(c *gin.Context) Remover {
	return NewGinCombineRemover(
		NewGinCookies(c, g.maxAge, g.secure),
	)
}
