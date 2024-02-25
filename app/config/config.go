package config

import "os"

var (
	// APIPort APIのポート
	APIPort string
	// RedisAddress Redisのアドレス
	RedisAddress string
)

func init() {
	APIPort = os.Getenv("API_PORT")
	RedisAddress = os.Getenv("REDIS_ADDRESS")
}
