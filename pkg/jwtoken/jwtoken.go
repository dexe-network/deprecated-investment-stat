package jwtoken

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

const (
	NoErr int = iota
	ErrMalformed
	ErrUnverifiable
	ErrSignature
	ErrExpired
	ErrClaims
	ErrCommon
)

type Account struct {
	ID           int64
	Role         string
	TokenExpires int64
}

type JWTConfig struct {
	Secret                 []byte
	AccessTokenExpiration  time.Duration
	RefreshTokenExpiration time.Duration
}

type JWT struct {
	log *zap.Logger
	cfg JWTConfig
}

func NewJWT(log *zap.Logger, cfg JWTConfig) *JWT {
	return &JWT{
		log: log,
		cfg: cfg,
	}
}

func (jt *JWT) GetAccessTokenExpiration() time.Duration {
	return jt.cfg.AccessTokenExpiration
}

func (jt *JWT) GetRefreshTokenExpiration() time.Duration {
	return jt.cfg.RefreshTokenExpiration
}

func (jt *JWT) ParseToken(t string) (account *Account, token *jwt.Token, errCode int, err error) {
	l := jt.log.With(zap.String("token", t))

	if token, err = jwt.Parse(t, jt.jwtKeyFunc); err != nil {
		l.Error("jwt.ParseToken(): ", zap.Error(err))
	}

	if err != nil {
		errVal := err.(*jwt.ValidationError)
		if errVal.Errors&(jwt.ValidationErrorMalformed) != 0 {
			errCode = ErrMalformed
		} else if errVal.Errors&(jwt.ValidationErrorUnverifiable) != 0 {
			errCode = ErrUnverifiable
		} else if errVal.Errors&(jwt.ValidationErrorSignatureInvalid) != 0 {
			errCode = ErrSignature
		} else if errVal.Errors&(jwt.ValidationErrorExpired) != 0 {
			errCode = ErrExpired
		} else if errVal.Errors&(jwt.ValidationErrorClaimsInvalid) != 0 {
			errCode = ErrClaims
		} else if errVal.Errors&(jwt.ValidationErrorAudience|jwt.ValidationErrorIssuedAt|jwt.ValidationErrorIssuer|jwt.ValidationErrorNotValidYet|jwt.ValidationErrorId) != 0 {
			errCode = ErrCommon
		}
	} else {
		errCode = NoErr
	}

	if errCode == ErrMalformed || errCode == ErrUnverifiable || errCode == ErrSignature || errCode == ErrClaims {
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		account = &Account{}
		account.TokenExpires = int64(claims["exp"].(float64))
		account.ID = int64(claims["sub"].(float64))
		account.Role = claims["rol"].(string)
		return
	}

	err = errors.New("cannot parse claims from a token")
	l.Error("token.Claims", zap.Error(err))
	return
}

func (jt *JWT) Parse(t string) (account Account, err error) {
	l := jt.log.With(zap.String("token", t))

	token, err := jwt.Parse(t, jt.jwtKeyFunc)
	if err != nil {
		l.Error("jwt.Parse", zap.Error(err))
		err = errors.New("cannot parse a token")
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		account.TokenExpires = int64(claims["exp"].(float64))
		account.ID = int64(claims["sub"].(float64))
		account.Role = claims["rol"].(string)
		return
	}

	err = errors.New("cannot parse claims from a token")
	l.Error("token.Claims", zap.Error(err))
	return
}

func (jt *JWT) jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return jt.cfg.Secret, nil
}

func (jt *JWT) Generate(account Account) (tokens Tokens, err error) {
	now := time.Now()
	iat := now.Unix()

	access := jwt.New(jwt.SigningMethodHS256)
	exp := now.Add(jt.cfg.AccessTokenExpiration).Unix()

	access.Claims = jwt.MapClaims{
		"exp": exp,
		"iat": iat,
		"sub": account.ID,
		"rol": account.Role,
	}

	if tokens.AccessToken, err = access.SignedString(jt.cfg.Secret); err != nil {
		return
	}

	refresh := jwt.New(jwt.SigningMethodHS256)
	exp = now.Add(jt.cfg.RefreshTokenExpiration).Unix()
	refresh.Claims = jwt.MapClaims{
		"exp": exp,
		"iat": iat,
		"sub": account.ID,
		"rol": account.Role,
	}

	if tokens.RefreshToken, err = refresh.SignedString(jt.cfg.Secret); err != nil {
		return
	}

	return
}

func (jt *JWT) GenerateForDuration(account Account, d time.Duration) (token string, err error) {
	now := time.Now()
	iat := now.Unix()

	access := jwt.New(jwt.SigningMethodHS256)
	exp := now.Add(d).Unix()

	access.Claims = jwt.MapClaims{
		"exp": exp,
		"iat": iat,
		"sub": account.ID,
		"rol": account.Role,
	}

	if token, err = access.SignedString(jt.cfg.Secret); err != nil {
		return
	}

	return
}
