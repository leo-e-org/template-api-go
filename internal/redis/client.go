package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"template-api-go/internal/logger"
)

var Client *redis.Client

func InitClient() {
	Client = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.host"),
		DB:       viper.GetInt("redis.db"),
		Password: viper.GetString("redis.password"),
	})

	if err := Ping(); err != nil {
		logger.Logger.Fatal("Failed to connect to Redis", zap.Error(err))
	}
}

func Ping() error {
	return Client.Ping(context.Background()).Err()
}
