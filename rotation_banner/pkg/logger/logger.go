package logger

import (
	"log"

	"go.uber.org/zap"
)

// Logger - алиас для быстрого доступа
var Logger *zap.SugaredLogger

func init() {
	var err error
	Logger, err = InitLogger()
	if err != nil {
		log.Fatal(err)
	}
}

// InitLogger - Инициализация логера
func InitLogger() (*zap.SugaredLogger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	Logger = logger.Sugar()
	return Logger, nil
}
