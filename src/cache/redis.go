package cache

import (
	"context"
	. "douyin/src/config"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()
var rc *redis.Client

func init() {
	host := AppConfig.GetString("redis.host")
	port := AppConfig.GetString("redis.port")
	password := AppConfig.GetString("redis.password")
	rc = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
}
func RCGet(key string) *redis.StringCmd {
	return rc.Get(ctx, key)
}
func RCExists(key string) bool {
	return rc.Exists(ctx, key).Val() != 0
}
func RCSet(key string, value interface{}, expiration time.Duration) {
	if RCExists(key) {
		rc.Expire(ctx, key, expiration)
		return
	}
	rc.Set(ctx, key, value, expiration)
}
func RCIncrement(key string) {
	rc.Incr(ctx, key)
}
func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}

func RCSAdd(key string, members interface{}) {
	rc.SAdd(ctx, key, members)
}

func RCSRem(key string, members interface{}) {
	rc.SRem(ctx, key, members)
}

func RCSmembers(key string) *redis.StringSliceCmd {
	return rc.SMembers(ctx, key)
}
