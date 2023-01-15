package log

import (
	"daimon/pkg/conf"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var L *zap.SugaredLogger

func InitLogger() {
	writeSyner := getWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyner, getLogLevel())
	logger := zap.New(core, zap.AddCaller())
	L = logger.Sugar()
}

func getLogLevel() zapcore.LevelEnabler {
	switch conf.C.App.Log {
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.DebugLevel
	}
}

func getEncoder() zapcore.Encoder {
	ec := zap.NewProductionEncoderConfig()
	if conf.C.App.Mode == "dev" {
		ec = zap.NewProductionEncoderConfig()
		ec.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		ec.EncodeLevel = zapcore.CapitalLevelEncoder
	}
	ec.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	return zapcore.NewConsoleEncoder(ec)
}

func getWriter() zapcore.WriteSyncer {

	lw := &lumberjack.Logger{
		Filename:   "logs/main.log",
		MaxSize:    100,
		MaxAge:     60,
		MaxBackups: 2,
	}

	if conf.C.App.LogStd {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lw), zapcore.AddSync(os.Stdout))
	} else {
		return zapcore.AddSync(lw)
	}
}
