package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

var loggerConf = zap.Config{
	OutputPaths:      []string{"stdout"},
	ErrorOutputPaths: []string{"stderr"},
	Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
	Encoding:         "console",
	EncoderConfig: zapcore.EncoderConfig{
		LevelKey:     "level",
		TimeKey:      "time",
		MessageKey:   "msg",
		CallerKey:    "caller",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	},
}

func SetupLogger(level string) {
	l, err := newLogger(level)
	if err != nil {
		log.Println(err)
	}
	Log = l
}

func newLogger(level string) (*zap.Logger, error) {
	conf := loggerConf
	conf.Level = zap.NewAtomicLevelAt(convertLevel(level))
	return conf.Build()
}
