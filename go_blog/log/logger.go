package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.SugaredLogger

var logLevel = zap.NewAtomicLevel()

func Setup() {

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:  "./my_goblog.txt",
		MaxSize:   1024, //MB
		LocalTime: true,
		Compress:  true,
	})

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		logLevel,
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	Log = logger.Sugar()
}

type Level int8

const (
	DebugLevel Level = iota - 1

	InfoLevel

	WarnLevel

	ErrorLevel

	DPanicLevel

	PanicLevel

	FatalLevel
)

func SetLevel(level Level) {
	logLevel.SetLevel(zapcore.Level(level))
}

func Info(args ...interface{}) {
	Log.Info(args...)
}

func Infof(template string, args ...interface{}) {
	Log.Infof(template, args...)
}

func Warn(args ...interface{}) {
	Log.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	Log.Warnf(template, args...)
}

func Error(args ...interface{}) {
	Log.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	Log.Errorf(template, args...)
}

func Panic(args ...interface{}) {
	Log.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	Log.Panicf(template, args...)
}
