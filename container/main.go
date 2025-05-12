package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	// 检查是否提供了 containerId 参数
	if len(os.Args) < 2 {
		fmt.Println("请提供 containerId 作为参数")
		return
	}
	containerId := os.Args[1]

	// 加载 .env 文件
	err := Load()
	if err != nil {
		fmt.Printf("加载 .env 文件时出错: %v\n", err)
		return
	}

	ContainerStats(containerId)
	// 你可以取消注释下面的函数调用以执行其他操作
	// ConatinerExec()
	// GetContainerLogsPro()
	// GetcContainer()
	// StopContainer()
	// Load()
	// everaiHome := os.Getenv("EVERAI_NODE_HOME")
	// fmt.Println("EVERAI_NODE_HOME:", everaiHome)
}

func Load() error {
	_, err := os.Stat(".env")
	if err == nil {
		return godotenv.Load()
	}

	if os.IsNotExist(err) {
		return nil
	}
	return err
}
