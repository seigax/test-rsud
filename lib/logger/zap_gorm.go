package logger

import (
	"context"
	"errors"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/otel"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

type ContextFn func(ctx context.Context) []zapcore.Field

type Logger struct {
	ZapLogger                 *zap.Logger
	LogLevel                  gormlogger.LogLevel
	SlowThreshold             time.Duration
	SkipCallerLookup          bool
	IgnoreRecordNotFoundError bool
	Context                   ContextFn
}

func newZapGorm(zapLogger *zap.Logger) Logger {
	return Logger{
		ZapLogger:                 zapLogger,
		LogLevel:                  gormlogger.Info,
		SlowThreshold:             100 * time.Millisecond,
		SkipCallerLookup:          false,
		IgnoreRecordNotFoundError: false,
		Context:                   nil,
	}
}

func (l Logger) SetAsDefault() {
	gormlogger.Default = l
}

func (l Logger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return Logger{
		ZapLogger:                 l.ZapLogger,
		SlowThreshold:             l.SlowThreshold,
		LogLevel:                  level,
		SkipCallerLookup:          l.SkipCallerLookup,
		IgnoreRecordNotFoundError: l.IgnoreRecordNotFoundError,
		Context:                   l.Context,
	}
}

func Info(ctx context.Context, str string, args ...interface{}) {
	zapGorm.Info(ctx, str, args)
}

func (l Logger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Info {
		return
	}
	l.logger(ctx).Sugar().Infof(str, args...)
	otel.Logger.InfoContext(ctx, str, args)
}

func Warn(ctx context.Context, str string, args ...interface{}) {
	zapGorm.Warn(ctx, str, args)
}

func (l Logger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Warn {
		return
	}
	l.logger(ctx).Sugar().Warnf(str, args...)
	otel.Logger.WarnContext(ctx, str, args)
}

func Error(ctx context.Context, str string, args ...interface{}) {
	zapGorm.Error(ctx, str, args)
}

func (l Logger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Error {
		return
	}
	l.logger(ctx).Sugar().Errorf(str, args...)
	otel.Logger.ErrorContext(ctx, str, args)
}

func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)
	logger := l.logger(ctx)
	sql, rows := fc()
	switch {
	case err != nil && l.LogLevel >= gormlogger.Error && (!l.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		logger.Error("trace", zap.Error(err), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
		otel.Logger.ErrorContext(ctx, "trace", zap.Error(err), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	case l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.LogLevel >= gormlogger.Warn:
		logger.Warn("trace", zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
		otel.Logger.WarnContext(ctx, "trace", zap.Error(err), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	case l.LogLevel >= gormlogger.Info:
		logger.Info("trace", zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
		otel.Logger.InfoContext(ctx, "trace", zap.Error(err), zap.Duration("elapsed", elapsed), zap.Int64("rows", rows), zap.String("sql", sql))
	}
}

var (
	gormPackage    = filepath.Join("gorm.io", "gorm")
	zapgormPackage = filepath.Join("moul.io", "zapgorm2")
)

func (l Logger) logger(ctx context.Context) *zap.Logger {
	logger := l.ZapLogger
	if l.Context != nil {
		fields := l.Context(ctx)
		logger = logger.With(fields...)
	}

	if l.SkipCallerLookup {
		return logger
	}

	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		case strings.Contains(file, gormPackage):
		case strings.Contains(file, zapgormPackage):
		default:
			return logger.WithOptions(zap.AddCallerSkip(i))
		}
	}
	return logger
}
