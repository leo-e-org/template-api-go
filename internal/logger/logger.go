package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger

func ecsEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// ECS layout fields
		TimeKey:       "@timestamp",
		LevelKey:      "@level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "@message",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
	}
}

func InitLogger() {
	ecsEncoder := zapcore.NewJSONEncoder(ecsEncoderConfig())
	ecsFileSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/api.json",
		MaxSize:    10,
		MaxBackups: 2,
		MaxAge:     1,
		Compress:   true,
	})

	consoleSyncer := zapcore.AddSync(os.Stdout)

	core := zapcore.NewTee(
		zapcore.NewCore(ecsEncoder, ecsFileSyncer, zapcore.InfoLevel),
		zapcore.NewCore(ecsEncoder, consoleSyncer, zapcore.DebugLevel),
	)

	Logger = zap.New(core)
	zap.ReplaceGlobals(Logger)

	if Logger == nil {
		panic("Failed to initialize logger")
	}
}

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		Logger.Info("Incoming request",
			zap.String("Method", c.Request.Method),
			zap.String("Path", c.Request.URL.Path),
			zap.String("IP", c.ClientIP()))
		c.Next()
	}
}
