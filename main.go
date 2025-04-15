package main

import (
	"go.uber.org/zap"
	"template-api-go/internal/api"
	"template-api-go/internal/config"
	"template-api-go/internal/logger"
	"template-api-go/internal/redis"
)

func main() {
	config.LoadConfig()
	logger.InitLogger()
	redis.InitClient()

	r := api.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		logger.Logger.Fatal("Server failed to start", zap.Error(err))
	}
}
