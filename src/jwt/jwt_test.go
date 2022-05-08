package jwt

import (
	"testing"
	"time"
)

func TestGenerateAndParseToken(t *testing.T) {

	token, err := GenerateToken()
	if err != nil {
		panic("Token生成错误")
	}
	println(token)

	myCustomClaims, err := ParseToken(token)

	println(myCustomClaims.ID)
	println(myCustomClaims.Username)
	println(myCustomClaims.ExpiresAt)
	println(time.Now().Unix())

}
