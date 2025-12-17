package logger

import (
	"contracts-manager/internal/infrastructure/config"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

func NewLogger(cfg *config.Config) *Logger {
	var zapLogger *zap.Logger
	var err error

	if cfg.RunMode == "dev" {
		zapLogger, err = zap.NewDevelopment()
	} else {
		conf := zap.NewProductionConfig()
		conf.EncoderConfig.TimeKey = "time"
		conf.EncoderConfig.MessageKey = "msg"
		conf.EncoderConfig.LevelKey = "level"
		conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		zapLogger, err = conf.Build()
	}

	if err != nil {
		log.Fatalf(ErrFailedInitLogger.Error()+": ", err)
	}

	return &Logger{zapLogger}
}

func (log *Logger) Errorf(msg error, err error) {
	log.Error(msg.Error(), zap.Error(err))
}
func (log *Logger) Fatalf(msg error, err error) {
	log.Fatal(msg.Error(), zap.Error(err))
}
