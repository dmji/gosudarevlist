package logger

import (
	"context"
	"fmt"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type loggerCtx string

const (
	loggerCtxValue loggerCtx = "logger"
)

type loggerZap struct {
	logger *zap.SugaredLogger
}

var globalLogger *loggerZap

func New() (*loggerZap, error) {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.Level.SetLevel(zapcore.DebugLevel)
	loggerConfig.ErrorOutputPaths = []string{"stdout"}
	//loggerConfig.EncoderConfig.MessageKey = zapcore.OmitKey

	logger, err := loggerConfig.Build(
		zap.AddCallerSkip(1),
	)
	if err != nil {
		return nil, fmt.Errorf("loggerConfig.Build: %w", err)
	}

	once := sync.Once{}
	once.Do(func() {
		globalLogger = &loggerZap{logger: logger.Sugar()}
	})

	return &loggerZap{logger: logger.Sugar()}, nil
}

func ToContext(ctx context.Context, logger *loggerZap) context.Context {
	return context.WithValue(ctx, loggerCtxValue, logger)
}

func FromContext(ctx context.Context) *loggerZap {
	if loggerC, ok := ctx.Value(loggerCtxValue).(*loggerZap); ok {
		return loggerC
	}

	return globalLogger
}

func Infow(ctx context.Context, msg string, keysAndValues ...interface{}) {
	if loggerC, ok := ctx.Value(loggerCtxValue).(*loggerZap); ok {
		loggerC.logger.Infow(msg, keysAndValues...)

		return
	}

	globalLogger.logger.Infow(msg, keysAndValues...)
}

func Errorw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	if loggerC, ok := ctx.Value(loggerCtxValue).(*loggerZap); ok {
		loggerC.logger.Errorw(msg, keysAndValues...)

		return
	}

	globalLogger.logger.Errorw(msg, keysAndValues...)
}

func Panicw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	if loggerC, ok := ctx.Value(loggerCtxValue).(*loggerZap); ok {
		loggerC.logger.Panicw(msg, keysAndValues...)

		return
	}

	globalLogger.logger.Panicw(msg, keysAndValues...)
}

func Fatalw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	if loggerC, ok := ctx.Value(loggerCtxValue).(*loggerZap); ok {
		loggerC.logger.Fatalw(msg, keysAndValues...)

		return
	}

	globalLogger.logger.Panicw(msg, keysAndValues...)
}
