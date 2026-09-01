package repository

import (
	"time"

	"github.com/phamvlap/url-shortener/internal/config"
	"github.com/phamvlap/url-shortener/internal/database"
	"github.com/redis/go-redis/v9"
)

type URLRepository struct {
	redisClient *redis.Client
}

func NewURLRepository() *URLRepository {
	// Get the configuration
	config := config.GetConfig()

	// Initialize the Redis client
	redisClient, err := database.GetRedisClient(config.Redis)
	if err != nil {
		panic(err)
	}

	return &URLRepository{
		redisClient: redisClient,
	}
}

func constructRedisKey(prefix string, params ...string) string {
	key := prefix
	for _, param := range params {
		key += ":" + param
	}
	return key
}

func (r *URLRepository) SaveURL(originalURL string, shortID string, expiration int64) error {
	expirationDuration := time.Duration(expiration) * time.Second
	return r.redisClient.Set(database.Ctx, constructRedisKey("url", shortID), originalURL, expirationDuration).Err()
}

func (r *URLRepository) GetOriginalURL(shortID string) (string, error) {
	return r.redisClient.Get(database.Ctx, constructRedisKey("url", shortID)).Result()
}

func (r *URLRepository) IncrCounter(shortID string) error {
	return r.redisClient.Incr(database.Ctx, constructRedisKey("counter", shortID)).Err()
}
