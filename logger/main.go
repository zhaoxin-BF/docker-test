package main

import (
	"fmt"
)

func main() {
	// 创建日志记录器
	tempDir := "/Users/zhaoxin/Workspace/github.com/zhaoxin-BF/docker-test/logger"
	fmt.Println("------------------: ", tempDir)
	log := NewLogger(tempDir)

	// 测试各种日志级别
	log.Debug("This is a debug message")
	log.Debugf("This is a debug message with format: %d", 123)
	log.Info("This is an info message")
	log.Infof("This is an info message with format: %s", "hello")
	log.Warn("This is a warning message")
	log.Warnf("This is a warning message with format: %.2f", 3.14)
	log.Error("This is an error message")
	log.Errorf("This is an error message with format: %v", struct{ Name string }{"John"})
}
