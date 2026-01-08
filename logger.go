package logger

import "go.uber.org/zap"

// 统一日志接口
type Logger interface {
	Debug(msg string, fields ...Field)
	Debugf(msg string, args ...interface{})
	Info(msg string, fields ...Field)
	Infof(msg string, args ...interface{})
	Warn(msg string, fields ...Field)
	Warnf(msg string, args ...interface{})
	Error(msg string, fields ...Field)
	Errorf(msg string, args ...interface{})

	With(fields ...Field) Logger
	Sync() error
}

// 日志字段类型
type Field = zap.Field

// 常用字段快捷方式
var (
	String   = zap.String
	Strings  = zap.Strings
	Int      = zap.Int
	Duration = zap.Duration
	Any      = zap.Any
	Err      = zap.NamedError
	Bool     = zap.Bool
	Int64    = zap.Int64
)

const (
	debug   = "debug"
	console = "console"
)
