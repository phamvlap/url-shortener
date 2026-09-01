package database

import (
	"context"
	"fmt"

	"sync"

	"github.com/phamvlap/url-shortener/internal/config"
	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var redisClient *redis.Client
var once sync.Once

func constructRedisURL(config config.RedisConfig) string {
	protocol := "redis"
	if config.TSLEnabled {
		protocol = "rediss"
	}
	return protocol + "://" + config.Username + ":" + config.Password + "@" + config.Host + ":" + config.Port + "/" + fmt.Sprintf("%d", config.DB)
}

func NewRedisClient(config config.RedisConfig) (*redis.Client, error) {
	redisURL := constructRedisURL(config)
	options, err := redis.ParseURL(redisURL)

	if err != nil {
		return nil, err
	}
	redisClient := redis.NewClient(options)

	if err := redisClient.Ping(Ctx).Err(); err != nil {
		redisClient.Close()
		return nil, err
	}

	return redisClient, nil
}

func GetRedisClient(config config.RedisConfig) (*redis.Client, error) {
	once.Do(func() {
		var err error
		redisClient, err = NewRedisClient(config)
		if err != nil {
			panic(err)
		}
	})
	return redisClient, nil
}
