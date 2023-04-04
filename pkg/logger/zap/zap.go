package zap

import (
	"context"

	"go.uber.org/zap"
)

type ZapLogger struct {
	lgr *zap.SugaredLogger
	ctx context.Context
}

func NewZapLogger(ctx context.Context) *ZapLogger {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	return &ZapLogger{lgr: sugar, ctx: ctx}
}

func (log *ZapLogger) Info(msg string) {
	log.lgr.Info(msg)
}

func (log *ZapLogger) Debug(msg string) {
	log.lgr.Info(msg)
}

func (log *ZapLogger) Panic(msg string) {
	log.lgr.Info(msg)
}

func (log *ZapLogger) Warn(msg string) {
	log.lgr.Info(msg)
}

func (log *ZapLogger) Error(msg string) {
	log.lgr.Info(msg)
}
