package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

func LokiTailLog() {
	// WebSocket 连接地址

	startTime := time.Now().Unix() - 60

	url := fmt.Sprintf("ws://192.168.31.43:30227/loki/api/v1/tail?query={app_id=\"GTmDE8iPXGSabeuU7qzN5G\",worker_id=\"EBK88vExFCNXQ5E5uqpBqZ\"}&start=%d", startTime)

	// 创建一个空的 HTTP 请求头
	header := http.Header{}

	// 建立 WebSocket 连接
	conn, _, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		log.Fatal("WebSocket 连接失败：", err)
	}
	defer conn.Close()

	// 接收并处理消息
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("读取消息错误：", err)
			return
		}
		fmt.Println("接收到消息：", string(message))
	}
	conn.Close()
}
