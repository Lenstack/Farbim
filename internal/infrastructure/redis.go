package infrastructure

import (
	"context"
	"github.com/go-redis/redis/v9"
	"log"
)

type Redis struct {
	host     string
	port     string
	password string
	db       int
	Client   *redis.Client
}

func NewRedisManager(host string, port string, password string) *Redis {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{Addr: host + ":" + port, Password: password, DB: 0})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("%s", err)
	}
	return &Redis{Client: rdb}
}
