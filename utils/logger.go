package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/Ilham-muttaqien17/learn-restful-go/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func getLogLevel() zapcore.Level {
	switch config.Env.GoEnv {
	case "production":
		return zapcore.InfoLevel
	case "development":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}

func NewLogger() *zap.Logger {
	// Define core list
	var cores []zapcore.Core

	// Define encoder config
	encoderConfig := zapcore.EncoderConfig{
		TimeKey: "timestamp",
		LevelKey: "level",
		MessageKey: "message",
		CallerKey: "caller",
		EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeLevel: zapcore.CapitalLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	// Define log levels
	logLevel := getLogLevel()

	// Create console encoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// Create cores for console
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), logLevel)

	// Append consoleCore into the core list
	cores = append(cores, consoleCore)

	isProduction := config.Env.GoEnv == "production"

	if isProduction {
		// Create file encoder
		fileEncoder := zapcore.NewJSONEncoder(encoderConfig)

		// Create lumberjack logger for file rotation
		logFileName := fmt.Sprintf("logs/error_%s.log", time.Now().Format("2006-01-02"))
		fileLogger := &lumberjack.Logger{
			Filename: logFileName,
			MaxSize: 10,
			MaxAge: 30,
			MaxBackups: 5,
			Compress: true,
		}

		// Create core for file
		fileCore := zapcore.NewCore(fileEncoder, zapcore.AddSync(fileLogger), logLevel)

		// Append consoleCore into the core list
		cores = append(cores, fileCore)
	}

	// Combine all cores
	core := zapcore.NewTee(cores...)

	// Create logger with caller info
	logger := zap.New(core, zap.AddCaller())

	Logger = logger

	// Ensure logs are flushed when the app exits
	defer logger.Sync()

	return logger
}