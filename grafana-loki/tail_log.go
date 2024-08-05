package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type LogData struct {
	Streams []Stream `json:"streams"`
}

type Stream struct {
	Stream map[string]string `json:"stream"`
	Values [][]string        `json:"values"`
}

func LokiTailLog() {
	// WebSocket 连接地址

	startTime := time.Now().Unix() - 60

	url := fmt.Sprintf("ws://192.168.31.43:30227/loki/api/v1/tail?query={app_id=\"XAjv5H4HVk7SbZJaNMjJhy\",worker_id=\"2VxrvwTdrTxn4gVQS56SuW\"}&start=%d", startTime)

	// 创建一个空的 HTTP 请求头
	header := http.Header{}

	// 建立 WebSocket 连接
	conn, _, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		log.Fatal("WebSocket 连接失败：", err)
	}
	defer conn.Close()

	go CloseOut(conn)

	// 接收并处理消息
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("读取消息错误：", err)
			return
		}

		res := &LogData{}

		err = json.Unmarshal(message, &res)
		if err != nil {
			fmt.Println("反序列化错误:", err)
			return
		}

		for _, stream := range res.Streams {
			fmt.Println(stream.Values[0][1])
		}
	}
	conn.Close()
}

func CloseOut(conn *websocket.Conn) {
	time.Sleep(100 * time.Second)
	err := conn.Close()
	if err != nil {
		fmt.Println(err)
	}
	err = conn.Close()
	if err != nil {
		fmt.Println(err)
	}
}
