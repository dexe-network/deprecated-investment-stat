package store

import (
	"net"
	"time"
)

type AccessToken struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	AccountID    int64     `json:"account_id"`
	ClientID     string    `json:"client_id"`
	UserAgent    string    `json:"string"`
	IP           net.IP    `json:"ip"`
	Expires      time.Time `json:"expires"`
	Created      time.Time `json:"created"`
	LastAccess   time.Time `json:"last_access"`
}
