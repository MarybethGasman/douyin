package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserTokenClaims struct {
	ID       int64
	Username string
	jwt.StandardClaims
}

func GenerateToken() (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(300 * time.Second)
	issuer := "frank"
	claims := UserTokenClaims{
		ID:       10001,
		Username: "frank",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("golang"))
	return token, err
}
func ParseToken(token string) (*UserTokenClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &UserTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("golang"), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*UserTokenClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
