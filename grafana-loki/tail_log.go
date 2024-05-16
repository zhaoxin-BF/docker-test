package grafana_loki

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func LokiTailLog() {
	// WebSocket 连接地址
	url := "ws://192.168.31.154:3100/loki/api/v1/tail?query={worker_id=\"worker-1\"}"

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
