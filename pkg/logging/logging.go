package logging

import (
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

type Config struct {
	Path     string
	Filename string
	Level    string
}

func NewLogger(config *Config) (*zap.Logger, error) {
	if _, err := os.Stat(config.Path); os.IsNotExist(err) {
		err := os.MkdirAll(config.Path, 0666)
		if err != nil {
			log.Fatalf("failed to make logs directory: %v", err)
		}
	}

	file, err := os.OpenFile(filepath.Join(config.Path, config.Filename),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("failed to open log file: %s, %v", config.Filename, err)
	}

	var level zapcore.Level
	switch config.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		log.Fatalf("invalid log level")
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "timestamp"
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	writer := zapcore.AddSync(file)
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, writer, level),
	)

	return zap.New(core, zap.AddCaller()), nil
}
