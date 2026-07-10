package config

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedis() (*redis.Client, error) {
	options, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(options)
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return client, nil
}