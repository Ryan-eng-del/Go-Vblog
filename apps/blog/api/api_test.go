package api

import (
	"github.com/infraboard/mcube/logger/zap"
	"testing"
)

func TestLog(t *testing.T) {
	// 获取全局Logger
	logger := zap.L()
	// 生成子Logger
	sub := logger.Named("sub.logger")

	sub.Debugf("this is a debug message, detail %s", "hello logger")
}

func init() {
	// 日志选项的手动设置
	// zap.Configure()
	// 直接使用默认配置
	// zap.DefaultConfig()

	// 直接调用开发模式zap.DevelopmentSetup(), 初始化开发者配置
	zap.DevelopmentSetup()
}
