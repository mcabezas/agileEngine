package logs

import (
	"context"

	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, args ...interface{})
	Info(v ...interface{})
	Infow(msg string, v ...interface{})
	Infof(format string, args ...interface{})
	Warn(v ...interface{})
	Warnf(format string, args ...interface{})
	Error(v ...interface{})
	Errorf(format string, args ...interface{})
	Errorw(msg string, v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, args ...interface{})
}

// InitDefault initializes a logger using uber-go/zap package in the application.
func NewSugaredLogger() Logger {
	conf := zap.Config{
		Encoding: "json",
		Level:    zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			CallerKey:    "file",
			MessageKey:   "msg",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		DisableCaller:    false,
	}
	log, _ := conf.Build()
	return log.Sugar()
}

func RequestID(ctx context.Context) zap.Field {
	return zap.String(RequestIDKey, middleware.GetReqID(ctx))
}
