package logger

import (
	"context"
	"sync"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	gormlogger "gorm.io/gorm/logger"
)

var (
	loggerOnce sync.Once
	zapGorm    Logger
)

func GetLogger() Logger {
	return zapGorm
}

func Init() {
	loggerOnce.Do(func() {
		zapConfig := zap.NewProductionConfig()
		zapConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		if viper.GetBool("LOG_USE_FILE") {
			zapConfig.OutputPaths = append(zapConfig.OutputPaths, "logs/app.log")
		}
		zapConfig.EncoderConfig.TimeKey = "@timestamp"
		zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
		logger, err := zapConfig.Build(zap.AddCallerSkip(1))
		if err != nil {
			panic(err)
		}

		zapGorm = newZapGorm(logger)
		zapGorm.SetAsDefault()
		zapGorm.LogLevel = gormlogger.Info
		zapGorm.SlowThreshold = time.Second

	})
}

func commonFields(ctx context.Context, message string, messages map[string]interface{}) []zap.Field {
	fields := []zap.Field{}

	requestID, ok := ctx.Value("X-Request-ID").(string)
	if ok {
		fields = append(fields, zap.String("request_id", requestID))
	}

	currentUserID, ok := ctx.Value("CurrentUserID").(uint)
	if ok {
		fields = append(fields, zap.Uint("actor_id", currentUserID))
	}

	for key, val := range messages {
		fields = append(fields, zap.Any(key, val))
	}

	return fields
}
