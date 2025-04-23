package main

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"template-api-go/internal/api"
	"template-api-go/internal/config"
	"template-api-go/internal/logger"
	"template-api-go/internal/redis"
	"time"
)

func main() {
	config.LoadConfig()
	logger.InitLogger()
	redis.InitClient()

	router := api.SetupRouter()
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	var wg = sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		logger.Logger.Info("Starting server", zap.String("address", server.Addr))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Logger.Fatal("Server failed to start", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Logger.Info("Shutdown signal received, shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Logger.Fatal("Server forced to shutdown", zap.Error(err))
	}

	if err := redis.CloseClient(); err != nil {
		logger.Logger.Error("Error closing Redis client connection", zap.Error(err))
	}

	wg.Wait()
	logger.Logger.Info("Server stopped")
}
