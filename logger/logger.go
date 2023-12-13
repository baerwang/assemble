package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

import (
	"os"
	"sync"
)

var logger Logger

// AsaLogger is logger struct
type AsaLogger struct {
	mutex sync.Mutex
	Logger
	dynamicLevel zap.AtomicLevel
	// disable presents the logger state. if disable is true, the logger will write nothing
	// the default value is false
	disable bool
}

// Logger
type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})
	Fatal(args ...interface{})

	Infof(fmt string, args ...interface{})
	Warnf(fmt string, args ...interface{})
	Errorf(fmt string, args ...interface{})
	Debugf(fmt string, args ...interface{})
	Fatalf(fmt string, args ...interface{})
}

func InitLogger(debug bool, level int, path string) {
	logger = NewLogger(debug, level, path)
}

func NewLogger(debug bool, level int, path string) *AsaLogger {
	var zapLoggerConfig zap.Config
	hook := &lumberjack.Logger{
		Filename:   path, // 日志文件路径
		MaxSize:    128,  // 每个日志文件保留的大小 单位:M
		MaxAge:     7,    // 文件最多保留多少天
		MaxBackups: 30,   // 日志文件最多保留多少个备份
		Compress:   true, // 是否压缩
	}
	defer hook.Close()

	zapLoggerConfig = zap.NewDevelopmentConfig()
	zapLoggerEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	zapLoggerConfig.EncoderConfig = zapLoggerEncoderConfig
	var writes = []zapcore.WriteSyncer{zapcore.AddSync(hook)}
	// 如果是开发环境，同时在管制台上也输入
	if debug {
		writes = append(writes, zapcore.AddSync(os.Stdout))
	}
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zapLoggerEncoderConfig), // 编码器配置
		zapcore.NewMultiWriteSyncer(writes...),            // 打印到控制台和文件
		zapcore.Level(level),                              // 日志等级
	)
	zapLogger := zap.New(core, zap.WithCaller(debug), zap.AddCallerSkip(1))
	return &AsaLogger{Logger: zapLogger.Sugar(), dynamicLevel: zapLoggerConfig.Level}
}

// Info
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Warn
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Error
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Debug
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Fatal
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Infof
func Infof(fmt string, args ...interface{}) {
	logger.Infof(fmt, args...)
}

// Warnf
func Warnf(fmt string, args ...interface{}) {
	logger.Warnf(fmt, args...)
}

// Errorf
func Errorf(fmt string, args ...interface{}) {
	logger.Errorf(fmt, args...)
}

// Debugf
func Debugf(fmt string, args ...interface{}) {
	logger.Debugf(fmt, args...)
}

// Fatalf
func Fatalf(fmt string, args ...interface{}) {
	logger.Fatalf(fmt, args...)
}
