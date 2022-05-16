package cache

import (
	"crypto/md5"
	"encoding/hex"
	"hash/crc32"
	"testing"
)

func TestRedisClient(t *testing.T) {
	t.Log(RCExists("ebba034e09d79a50692371d0070e61e"))

}

func CRC32(input string) uint32 {
	bytes := []byte(input)
	return crc32.ChecksumIEEE(bytes)
}
func MD5(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}
