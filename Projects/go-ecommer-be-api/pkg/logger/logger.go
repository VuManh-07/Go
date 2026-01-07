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
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level,
	)

	// logger := zap.New(core, zap.AddCaller())
	// zap.ReplaceGlobals(logger)
	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

// format log as message
func getEncodeLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}
