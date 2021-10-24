package logger

import (
	"context"

	"go.uber.org/zap"
)

type key int

var ctxKey key

func GetLoggerFromContext(ctx context.Context) *zap.Logger {
	return ctx.Value(ctxKey).(*zap.Logger)
}

func SetLoggerInContext(ctx context.Context, log *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxKey, log)
}