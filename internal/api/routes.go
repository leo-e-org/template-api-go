package api

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"go.uber.org/zap"
	"template-api-go/internal/controller"
	"template-api-go/internal/logger"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	if err := router.SetTrustedProxies([]string{}); err != nil {
		logger.Logger.Fatal("Error setting proxy configuration", zap.Error(err))
	}

	router.Use(gin.Recovery())
	router.Use(logger.GinLogger())

	prometheus := ginprometheus.NewPrometheus("go_api")
	prometheus.MetricsPath = "/template-api-go/metrics"

	api := router.Group("/template-api-go")
	api.GET("/app-version", controller.GetVersion)

	api.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.GET(prometheus.MetricsPath, gin.WrapH(promhttp.Handler()))

	return router
}
