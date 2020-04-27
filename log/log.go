package log

import (
	"flag"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	debug  bool
	Suger  *zap.SugaredLogger
	Logger *zap.Logger
)

const (
	loggerMaxSize    = 100
	loggerMaxBackups = 3
	loggerMaxAge     = 9
)

func init() {
	flag.BoolVar(&debug, "debug", false, "log debug")
	flag.Parse()
	var (
		encoder zapcore.Encoder
		level   zapcore.Level
	)
	if debug {
		encoder = zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
		level = zap.DebugLevel
	} else {
		encoder = zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
		level = zap.InfoLevel
	}

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "history.log",
		MaxSize:    loggerMaxSize, // megabytes
		MaxBackups: loggerMaxBackups,
		MaxAge:     loggerMaxAge, // days
	})
	core := zapcore.NewCore(
		encoder,
		w,
		level,
	)
	Logger = zap.New(core, zap.AddCaller())
	defer Logger.Sync()
	Suger = Logger.Sugar()
}
