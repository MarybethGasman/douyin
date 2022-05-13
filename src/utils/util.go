package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func MD5WithSalt(s string) string {
	sum := md5.Sum([]byte(s + "sandaiyidui"))
	return hex.EncodeToString(sum[:])
}
