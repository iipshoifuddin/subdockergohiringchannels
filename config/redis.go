package config

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

type redisClient struct {
	c *redis.Client
}

var (
	client = &redisClient{}
)

//GetClient get the redis client
func SetupRedis() *redis.Client {
	addr := fmt.Sprintf(os.Getenv("REDIS_CONN_STRING"))
	pass := fmt.Sprintf(os.Getenv("REDIS_CONN_PASS_STRING"))
	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
	})

	if err := c.Ping().Err(); err != nil {
		panic("Unable to connect to redis " + err.Error())
	}
	return c
}
