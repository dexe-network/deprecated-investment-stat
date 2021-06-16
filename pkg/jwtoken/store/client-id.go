package store

import (
	"go.uber.org/zap"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	CookieClientID = "client-id"
)

type ClientID struct {
	id           string
	log          *zap.Logger
	ctx          *gin.Context
	cookieSecure bool
	cookieMaxAge int
}

func NewClientID(
	ctx *gin.Context,
	log *zap.Logger,
	cookieSecure bool,
	cookieMaxAge time.Duration,
) (c ClientID) {
	c = ClientID{
		ctx:          ctx,
		log:          log,
		cookieSecure: cookieSecure,
		cookieMaxAge: int(cookieMaxAge / time.Second),
	}
	c.ReRead()
	return
}

func (o *ClientID) Generate() (err error) {
	var id uuid.UUID
	if id, err = uuid.NewRandom(); err != nil {
		return
	}
	o.id = id.String()
	return
}

func (o *ClientID) Write(ctx *gin.Context) {
	if o.id == "" {
		o.log.Warn("jwttoken.store.ClientID trying to set an empty client-id cookie")
	}
	ctx.SetCookie(CookieClientID, o.id, o.cookieMaxAge, "/", "", o.cookieSecure, true)
}

func (o *ClientID) ReRead() {
	c, err := o.ctx.Cookie(CookieClientID)
	if err != nil || c == "" {
		o.id = ""
		return
	}
	o.id = c
}

func (o *ClientID) Get() string {
	return o.id
}
