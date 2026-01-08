package logger

import (
	"fmt"
	"os"

	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	Messagekey = "message"
	LevelKey   = "level"
)

var (
	once sync.Once
)

//实现zap日志器的封装

// ZapLogger zap 实现
type ZapLogger struct {
	logger *zap.Logger
	sugar  *zap.SugaredLogger
	config LogConfig
}

// NewZapLogger 创建新的zapLogger 实例
func NewZapLogger(cfg LogConfig) (Logger, error) {
	//设置默认值
	cfg.SetDefault()

	//创建核心列表
	cores := []zapcore.Core{}

	//设置日志级别,创建可以动态变更的日志级别对象
	level := zap.NewAtomicLevel()
	if err := level.UnmarshalText([]byte(cfg.Level)); err != nil {
		return nil, err
	}

	//编码器配置
	var encordeCfg zapcore.EncoderConfig
	if cfg.Devolpment {
		encordeCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encordeCfg = zap.NewProductionEncoderConfig()
	}

	//时间格式化
	encordeCfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	//将日志输出info改为INFO
	encordeCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	encordeCfg.MessageKey = Messagekey
	encordeCfg.LevelKey = LevelKey

	//控制台输出
	if cfg.EnableConsole {
		encordeCfg.EncodeLevel = ColorLevelEncoder
		consoleEncoder := zapcore.NewConsoleEncoder(encordeCfg)
		consoleCore := zapcore.NewCore(consoleEncoder,
			zapcore.Lock(os.Stdout),
			level,
		)
		cores = append(cores, consoleCore)
	}
	//文件输出
	if cfg.FileConsole {
		// 创建文件专用的编码器配置（不带颜色）
		fileEncCfg := zapcore.EncoderConfig{
			TimeKey:        "T",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder, // 文件不使用颜色
			EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}

		FileEncoder := zapcore.NewJSONEncoder(fileEncCfg)
		FileWiter := zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.Directory + "/" + cfg.FileName,
			MaxSize:    int(cfg.MaxSize),
			MaxBackups: int(cfg.MaxBackups),
			MaxAge:     int(cfg.MaxAge),
			Compress:   cfg.Compress,
			LocalTime:  true,
		})
		FileCore := zapcore.NewCore(FileEncoder, FileWiter, level)
		cores = append(cores, FileCore)
	}
	//创建核心
	core := zapcore.NewTee(cores...)

	//创建日志器
	options := []zap.Option{
		//当你调用 logger.Error(...) 时，日志输出会包含调用栈。
		zap.AddStacktrace(zap.ErrorLevel),
	}
	if cfg.EnableCaller {
		options = append(options, zap.AddCaller())
	}
	//跳过文件调用层数
	skilloption := zap.AddCallerSkip(2)
	options = append(options, skilloption)

	logger := zap.New(core, options...)
	sugar := logger.Sugar()

	return &ZapLogger{
		logger: logger,
		sugar:  sugar,
		config: cfg,
	}, nil
}

// 实现Logger接口
func (l *ZapLogger) Debug(msg string, fields ...Field) {
	l.logger.Debug(msg, fields...)
}
func (l *ZapLogger) Info(msg string, fields ...Field) {
	l.logger.Info(msg, fields...)
}
func (l *ZapLogger) Warn(msg string, fields ...Field) {
	l.logger.Warn(msg, fields...)
}

func (l *ZapLogger) Error(msg string, fields ...Field) {
	l.logger.Error(msg, fields...)
}

func (l *ZapLogger) Debugf(msg string, args ...interface{}) {
	l.sugar.Debugf(msg, args...)
}

func (l *ZapLogger) Infof(msg string, args ...interface{}) {
	l.sugar.Infof(msg, args...)
}

func (l *ZapLogger) Warnf(msg string, args ...interface{}) {
	l.sugar.Warnf(msg, args...)
}

func (l *ZapLogger) Errorf(msg string, args ...interface{}) {
	l.sugar.Errorf(msg, args...)
}

func (l *ZapLogger) With(fields ...Field) Logger {
	return &ZapLogger{
		logger: l.logger.With(fields...),
		sugar:  l.sugar.With(fieldsToInterfaces(fields)),
		config: l.config,
	}
}

func (l *ZapLogger) Sync() error {
	return l.logger.Sync()
}
func fieldsToInterfaces(fields []Field) []interface{} {
	res := make([]interface{}, len(fields))
	for i, f := range fields {
		res[i] = f
	}
	return res
}

// InitGlobalLogger 初始化全局日志
func InitGlobalLogger(config LogConfig) error {
	var err error
	once.Do(func() {
		var l Logger
		l, err = NewZapLogger(config)
		if err == nil {
			SetGlobalLogger(l)
		}
	})
	if err != nil {
		return fmt.Errorf("failed to init logger %v", err)
	}
	return nil
}
