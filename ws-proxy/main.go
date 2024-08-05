package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	// 定义 WebSocket 升级器
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // 允许所有源访问
		},
	}

	// 创建 HTTP 处理程序
	http.HandleFunc("/api/loggers/v1/tail", func(w http.ResponseWriter, r *http.Request) {
		appID := r.URL.Query().Get("app_id")
		workerID := r.URL.Query().Get("worker_id")

		fmt.Println(appID)
		fmt.Println(workerID)
		// 从客户端升级 WebSocket 连接
		clientConn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Failed to upgrade WebSocket connection: %v", err)
			return
		}
		defer clientConn.Close()

		// 连接到远程 WebSocket 服务器
		//u := url.URL{Scheme: "ws", Host: "127.0.0.1:8080", Path: "/ws"}
		startTime := time.Now().Unix() - 60

		url := fmt.Sprintf("ws://192.168.31.43:30227/loki/api/v1/tail?query={app_id=\"XAjv5H4HVk7SbZJaNMjJhy\",worker_id=\"2VxrvwTdrTxn4gVQS56SuW\"}&start=%d", startTime)
		remoteConn, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			log.Printf("Failed to connect to remote WebSocket server: %v", err)
			return
		}
		defer remoteConn.Close()

		// 在客户端和远程服务器之间转发数据
		go copyData(clientConn, remoteConn)
		copyData(remoteConn, clientConn)
	})

	fmt.Println("Starting WebSocket proxy server on :8088")
	log.Fatal(http.ListenAndServe(":8088", nil))
}

func copyData(dst, src *websocket.Conn) {
	for {
		messageType, message, err := src.ReadMessage()
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Printf("Failed to read message: %v", err)
			return
		}

		if err := dst.WriteMessage(messageType, message); err != nil {
			log.Printf("Failed to write message: %v", err)
			return
		}
	}
}
