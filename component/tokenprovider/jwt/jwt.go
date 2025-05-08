package jwt

import (
	tokenProvider "g05-food-delivery/component/tokenprovider"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtProvider struct {
	secret string
}

func NewTokenJwtProvider(secret string) *JwtProvider {
	return &JwtProvider{
		secret: secret,
	}
}

type myClaims struct {
	Payload tokenProvider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (j *JwtProvider) Generate(data tokenProvider.TokenPayload, expire int) (*tokenProvider.Token, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		Payload: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(expire)).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	return &tokenProvider.Token{
		Token:   myToken,
		Expiry:  expire,
		Created: time.Now(),
	}, nil
}

func (j *JwtProvider) Validate(token string) (*tokenProvider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(token, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, tokenProvider.ErrInvalidToken
	}

	if !res.Valid {
		return nil, tokenProvider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, tokenProvider.ErrInvalidToken
	}

	return &claims.Payload, nil
}

func (j *JwtProvider) String() string {
	return "JWT implement Provider"
}
