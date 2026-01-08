package logger

import (
	"errors"
	"testing"
)

func TestNewZapLogger(t *testing.T) {
	cfg := LogConfig{
		Level:         debug,
		EnableConsole: true,
		EnableCaller:  true,
	}
	l, err := NewZapLogger(cfg)
	if err != nil {
		t.Fatalf("Failed to create ZapLogger: %v", err)
	}
	if l == nil {
		t.Fatal("Logger instance is nil")
	}
}

func TestLoggerMethods(t *testing.T) {
	cfg := LogConfig{
		Level:         debug,
		EnableConsole: true,
		EnableCaller:  true,
	}
	l, _ := NewZapLogger(cfg)

	// 不检查输出内容，只要不panic即可
	l.Debug("debug msg")
	l.Info("info msg")
	l.Warn("warn msg")
	l.Error("error msg")
	l.Debugf("%s msg", "debugf")
	l.Infof("%s msg", "infof")
	l.Warnf("%s msg", "warnf")
	l.Errorf("%s msg", "errorf")

	// 测试With
	l2 := l.With(String("foo", "bar"), Err("TEST", errors.New("test")))
	l2.Info("with msg")

	// 测试Sync
	if err := l.Sync(); err != nil {
		t.Logf("Sync error: %v", err)
	}
}

func TestGlobalLoggerMethods(t *testing.T) {
	// 直接调用全局方法
	Info("global info", String("k", "v"))
	Warn("global warn")
	Error("global error")
	WithFields(String("k1", "v1")).Info("with fields")
	if err := Sync(); err != nil {
		t.Logf("Global Sync error: %v", err)
	}
}
