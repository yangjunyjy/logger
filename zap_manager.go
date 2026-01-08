package logger

import (
	"go.uber.org/zap"
)

var (
	globalLogger Logger
)

func SetGlobalLogger(l Logger) {
	globalLogger = l
}

// L 获取全局日志实例
func L() Logger {
	if globalLogger == nil {
		// 创建默认日志器作为后备
		defaultConfig := LogConfig{
			Level:         debug,
			Encoding:      console,
			EnableConsole: true,
			EnableCaller:  true,
		}
		globalLogger, _ = NewZapLogger(defaultConfig)
	}
	return globalLogger
}

func Debug(msg string, fields ...zap.Field) {
	L().Debug(msg, fields...)
}

func Debugf(msg string, args ...interface{}) {
	L().Debugf(msg, args...)
}

// Info 全局Info日志
func Info(msg string, fields ...zap.Field) {
	L().Info(msg, fields...)
}

func Infof(msg string, args ...interface{}) {
	L().Infof(msg, args...)
}

// Warn 全局Warn日志
func Warn(msg string, fields ...zap.Field) {
	L().Warn(msg, fields...)
}

// Warn 全局Warn日志
func Warnf(msg string, args ...interface{}) {
	L().Warnf(msg, args...)
}

// Error 全局Error日志
func Error(msg string, fields ...zap.Field) {
	L().Error(msg, fields...)
}

func Errorf(msg string, args ...interface{}) {
	L().Errorf(msg, args...)
}

// WithFields 创建带字段的日志器
func WithFields(fields ...zap.Field) Logger {
	return L().With(fields...)
}

// Sync 刷新日志缓冲区
func Sync() error {
	return L().Sync()
}
