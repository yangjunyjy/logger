package logger

import (
	"fmt"

	"go.uber.org/zap/zapcore"
)

// ANSI颜色码
var (
	levelColors = map[zapcore.Level]string{
		zapcore.DebugLevel:  "\033[36m", // 青色
		zapcore.InfoLevel:   "\033[32m", // 绿色
		zapcore.WarnLevel:   "\033[33m", // 黄色
		zapcore.ErrorLevel:  "\033[31m", // 红色
		zapcore.DPanicLevel: "\033[35m", // 紫色
		zapcore.PanicLevel:  "\033[35m", // 紫色
		zapcore.FatalLevel:  "\033[35m", // 紫色
	}
	colorReset = "\033[0m"
)

// 自定义LevelEncoder
func ColorLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	color, ok := levelColors[level]
	if !ok {
		color = colorReset
	}
	enc.AppendString(fmt.Sprintf("%s%s%s", color, level.CapitalString(), colorReset))
}
