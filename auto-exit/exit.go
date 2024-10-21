package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main01() {
	// 发送 SIGINT 信号给自身进程
	go func() {
		time.Sleep(10 * time.Second)
		pid := os.Getpid()
		err := syscall.Kill(pid, syscall.SIGINT)
		if err != nil {
			fmt.Println("发送 SIGINT 信号失败:", err)
			return
		}
	}()

	// 创建信号监听通道
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// 等待信号
	fmt.Println("程序正在运行,等待 SIGINT 信号...")
	sig := <-signalChan
	fmt.Printf("收到信号: %v\n", sig)

	// 执行清理操作
	cleanup()

	// 退出程序
	os.Exit(0)
}

func cleanup() {
	// 在这里添加程序下线前需要执行的清理操作
	fmt.Println("正在执行清理操作...")
}
