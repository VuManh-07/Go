package logger

import (
	"os"

	"github.com/VuManh-07/Go/Projects/go-ecommer-be-api/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(loggerConfig setting.LoggerConfig) *LoggerZap {
	logLevel := loggerConfig.Level
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	encode := getEncodeLog()
	hook := lumberjack.Logger{
		Filename:   loggerConfig.Filename,
		MaxSize:    loggerConfig.MaxSize,    // megabytes
		MaxBackups: loggerConfig.MaxBackups, // 5files
		MaxAge:     loggerConfig.MaxAge,     // days
		Compress:   true,
	}
	core := zapcore.NewCore(
		encode,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // write to file and console
		level,
	)

	return &LoggerZap{
		zap.New(core, zap.AddCaller(), // show line number
			zap.AddStacktrace(zap.ErrorLevel)), // show stacktrace when level >= error
	}
}

// format log as message
func getEncodeLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // format time 2024-05-20T15:04:05.000Z0700
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder // format level INFO, ERROR

	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}
