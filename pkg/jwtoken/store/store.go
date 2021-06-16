package store

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net"
	"time"

	"dex-trades-parser/pkg/jwtoken"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type Store struct {
	log         *zap.Logger
	redisClient redis.Cmdable
	redisPrefix string
	jwt         *jwtoken.JWT
}

func New(
	log *zap.Logger,
	redisClient redis.Cmdable,
	redisPrefix string,
	jwt *jwtoken.JWT,
) *Store {
	return &Store{
		log:         log,
		redisClient: redisClient,
		redisPrefix: redisPrefix,
		jwt:         jwt,
	}
}

func (o *Store) newAccessTokenStruct(account jwtoken.Account, tokens jwtoken.Tokens, expires time.Duration, clientID string) AccessToken {
	return AccessToken{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		AccountID:    account.ID,
		UserAgent:    "Mozilla/Firefox",      // @todo: set real user agent
		IP:           net.IPv4(127, 0, 0, 1), // @todo: set real user IP address, to show user info about location of login
		Expires:      time.Now().Add(expires),
		Created:      time.Now(),
		LastAccess:   time.Now(),
		ClientID:     clientID,
	}
}

func (o *Store) accessTokenAllKeysPattern(accountID int64) string {
	return fmt.Sprintf("%s:authorization:%d:*", o.redisPrefix, accountID)
}

func (o *Store) accessTokenAllForClientKeysPattern(clientID string, accountID int64) string {
	return fmt.Sprintf("%s:authorization:%d:%s:*", o.redisPrefix, accountID, clientID)
}

func (o *Store) accessTokenKey(at, clientID string, accountID int64) string {
	return fmt.Sprintf("%s:authorization:%d:%s:%x", o.redisPrefix, accountID, clientID, sha256.Sum256([]byte(at)))
}

func (o *Store) accessTokenFromStorage(at, clientID string, accountID int64) (ats AccessToken, err error) {
	var b []byte
	if b, err = o.redisClient.Get(o.accessTokenKey(at, clientID, accountID)).Bytes(); err != nil {
		return
	}

	ats = AccessToken{}
	if err = json.Unmarshal(b, &ats); err != nil {
		return
	}

	return
}

func (o *Store) currentTokens(atString, clientID string, accountID int64) (at AccessToken, err error) {
	if at, err = o.accessTokenFromStorage(atString, clientID, accountID); err != nil {
		return
	}
	return
}

func (o *Store) Logout(accountID int64, currentATString string, allButCurrent bool, cid ClientID, ctx *gin.Context) (err error) {
	if cid.Get() == "" {
		if err = cid.Generate(); err != nil {
			return
		}
		cid.Write(ctx)
	}

	var currentAT AccessToken
	if currentAT, err = o.accessTokenFromStorage(currentATString, cid.Get(), accountID); err != nil {
		return
	}

	if allButCurrent {
		// logout all sessions, excluding current
		var keys []string
		atKey := o.accessTokenKey(currentAT.AccessToken, cid.Get(), accountID)
		if keys, err = o.redisClient.Keys(o.accessTokenAllKeysPattern(accountID)).Result(); err != nil {
			return
		}
		for _, v := range keys {
			if v == atKey {
				continue
			}
			if e := o.redisClient.Del(v).Err(); e != nil {
				o.log.Error("failed to delete access token data while logout", zap.Error(e), zap.String("value", v))
			}
		}
	} else {
		atKey := o.accessTokenKey(currentAT.AccessToken, cid.Get(), accountID)
		if err = o.redisClient.Del(atKey).Err(); err != nil {
			return
		}
	}

	return
}

func (o *Store) Login(account jwtoken.Account, cid ClientID, ctx *gin.Context) (tokens jwtoken.Tokens, err error) {
	if cid.Get() == "" {
		if err = cid.Generate(); err != nil {
			return
		}
		cid.Write(ctx)
	}

	if tokens, err = o.jwt.Generate(account); err != nil {
		return
	}

	expires := o.jwt.GetRefreshTokenExpiration()

	var atb []byte
	at := o.newAccessTokenStruct(account, tokens, expires, cid.Get())
	if atb, err = json.Marshal(at); err != nil {
		o.log.Error("failed to encode access token data before storing to redis", zap.Error(err))
		return
	}
	atKey := o.accessTokenKey(tokens.AccessToken, cid.Get(), account.ID)
	if err = o.redisClient.Set(atKey, atb, expires).Err(); err != nil {
		o.log.Error("failed to store access token to redis: %s", zap.Error(err))
		return
	}

	return
}

func (o *Store) ExtendAccessToken(currentATString string, cid ClientID, ctx *gin.Context) (newTokens jwtoken.Tokens, prevTokens jwtoken.Tokens, errCode int, err error) {
	if cid.Get() == "" {
		if err = cid.Generate(); err != nil {
			return
		}
		cid.Write(ctx)
	}

	var at AccessToken
	var account *jwtoken.Account

	errCode = jwtoken.NoErr

	account, _, errCode, err = o.jwt.ParseToken(currentATString)
	if err == nil {
		at, err = o.currentTokens(currentATString, cid.Get(), account.ID)
		if err != nil {
			return
		}
		prevTokens = jwtoken.Tokens{AccessToken: at.AccessToken, RefreshToken: at.RefreshToken}
		newTokens = prevTokens
		return
	}

	if account == nil {
		return
	}

	at, err = o.currentTokens(currentATString, cid.Get(), account.ID)
	if err != nil {
		return
	}
	prevTokens = jwtoken.Tokens{AccessToken: at.AccessToken, RefreshToken: at.RefreshToken}

	if newTokens, err = o.jwt.Generate(*account); err != nil {
		return
	}

	expires := o.jwt.GetRefreshTokenExpiration()

	var atNewB []byte
	atNew := o.newAccessTokenStruct(*account, newTokens, expires, cid.Get())
	atNew.Created = at.Created

	if atNewB, err = json.Marshal(at); err != nil {
		o.log.Error("failed to encode access token data before storing to redis", zap.Error(err))
		return
	}
	atNewKey := o.accessTokenKey(newTokens.AccessToken, cid.Get(), account.ID)
	if err = o.redisClient.Set(atNewKey, atNewB, expires).Err(); err != nil {
		o.log.Error("failed to store access token to redis", zap.Error(err))
		return
	}

	// remove old tokens
	var keys []string
	if keys, err = o.redisClient.Keys(o.accessTokenAllForClientKeysPattern(cid.Get(), account.ID)).Result(); err != nil {
		return
	}
	for _, v := range keys {
		if v == atNewKey {
			continue
		}
		if e := o.redisClient.Expire(v, time.Duration(time.Minute)).Err(); e != nil {
			o.log.Error("failed to delete access token data while logout", zap.Error(err), zap.String("value", v))

		}
	}

	return
}
