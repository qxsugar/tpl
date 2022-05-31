package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type LoginRole int

const (
	Header = "Token"
	UID    = "uid"
)

var (
	SecretKey      = "api-project"
	ExpireDuration = time.Minute * 60 * 24 * 7
)

type AuthClaims struct {
	jwt.StandardClaims
	Uid int
}

func NewToken(uid int) (string, error) {
	expire := time.Now().Add(ExpireDuration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, AuthClaims{
		Uid:            uid,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expire.Unix()},
	})
	return token.SignedString([]byte(SecretKey))
}

func ToUid(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.Uid, nil
	}
	return 0, errors.New("token is valid")
}
