package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	App   AppConfig
	Redis RedisConfig
}

type AppConfig struct {
	Port     string
	Domain   string
	APIQuota int
}

type RedisConfig struct {
	Host       string
	Port       string
	Password   string
	DB         int
	Username   string
	TSLEnabled bool
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

func getEnvAsInt(name string, defaultValue int) int {
	valueStr := os.Getenv(name)
	if valueStr == "" {
		return defaultValue
	}

	var value int
	_, err := fmt.Sscanf(valueStr, "%d", &value)
	if err != nil {
		return defaultValue
	}

	return value
}

func getEnvAsBool(name string, defaultValue bool) bool {
	valueStr := os.Getenv(name)
	if valueStr == "" {
		return defaultValue
	}

	var value bool
	_, err := fmt.Sscanf(valueStr, "%t", &value)
	if err != nil {
		return defaultValue
	}

	return value
}

func Load() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	return Config{
		App: AppConfig{
			Port:     getEnv("API_PORT", "8000"),
			Domain:   getEnv("DOMAIN", "localhost"),
			APIQuota: getEnvAsInt("API_QUOTA", 1000),
		},
		Redis: RedisConfig{
			Host:       getEnv("REDIS_HOST", "localhost"),
			Port:       getEnv("REDIS_PORT", "6379"),
			Password:   getEnv("REDIS_PASSWORD", ""),
			DB:         getEnvAsInt("REDIS_DB", 0),
			Username:   getEnv("REDIS_USERNAME", ""),
			TSLEnabled: getEnvAsBool("REDIS_TLS_ENABLED", false),
		},
	}, nil
}

var config Config
var once sync.Once

func GetConfig() Config {
	once.Do(func() {
		var err error
		config, err = Load()
		if err != nil {
			panic(err)
		}
	})
	return config
}
