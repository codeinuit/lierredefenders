package main

import (
	"context"

	"github.com/codeinuit/lierredefenders/pkg/logger"
	"github.com/codeinuit/lierredefenders/pkg/logger/zap"
)

type LierreBot struct {
	logger logger.Logger
}

func main() {
	bot := LierreBot{
		logger: zap.NewZapLogger(context.TODO()),
	}

	bot.logger.Info("test")
}
