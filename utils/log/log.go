package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *Logger

func init() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	switch os.Getenv("TCMS_LOG_LEVEL") {
	case "debug":
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warning":
		cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "panic":
		cfg.Level = zap.NewAtomicLevelAt(zap.PanicLevel)
	case "fatal":
		cfg.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}
	logger, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	Log = &Logger{logger.Sugar()}
}

type Logger struct {
	Sugar *zap.SugaredLogger
}

func (l *Logger) Debugf(msg string, args ...interface{}) {
	l.Sugar.Debugf(msg, args...)
}

func (l *Logger) Infof(msg string, args ...interface{}) {
	l.Sugar.Infof(msg, args...)
}

func (l *Logger) Warnf(msg string, args ...interface{}) {
	l.Sugar.Warnf(msg, args...)
}

func (l *Logger) Errorf(err error, msg string, args ...interface{}) {
	l.Sugar.With("err", err).Errorf(msg, args...)
}

func (l *Logger) Fatalf(msg string, args ...interface{}) {
	l.Sugar.Fatalf(msg, args...)
}

func (l *Logger) With(args ...interface{}) *Logger {
	return &Logger{l.Sugar.With(args...)}
}

func (l *Logger) Desugar() *zap.Logger {
	return l.Sugar.Desugar().WithOptions(zap.AddCallerSkip(-1))
}

func Debugf(msg string, args ...interface{}) {
	Log.Sugar.Debugf(msg, args...)
}

func Infof(msg string, args ...interface{}) {
	Log.Sugar.Infof(msg, args...)
}

func Warnf(msg string, args ...interface{}) {
	Log.Sugar.Warnf(msg, args...)
}

func Errorf(err error, msg string, args ...interface{}) {
	Log.Sugar.With("err", err).Errorf(msg, args...)
}

func Fatalf(msg string, args ...interface{}) {
	Log.Sugar.Fatalf(msg, args...)
}

func With(args ...interface{}) *Logger {
	return &Logger{Log.Sugar.With(args...)}
}
