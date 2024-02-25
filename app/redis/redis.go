package redis

import (
	"github.com/redis/go-redis/v9"
	"github.com/wasuwa/blog-api/app/config"
)

// NewClient redisのクライアントを初期化します
func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: config.RedisAddress,
	})
	return client
}
