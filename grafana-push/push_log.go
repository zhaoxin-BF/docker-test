package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type StreamData struct {
	Streams []Stream `json:"streams"`
}

type Stream struct {
	StreamInfo StreamInfo      `json:"stream"`
	Values     [][]interface{} `json:"values"`
}

type StreamInfo struct {
	WorkerId    string `json:"worker_id"`
	ServiceName string `json:"service_name"`
}

func main() {
	url := "http://192.168.31.154:3100/loki/api/v1/push"

	// 构建请求体数据
	currentTime := time.Now().UnixNano()
	timeString := strconv.FormatInt(currentTime, 10)
	data := StreamData{
		Streams: []Stream{
			{
				StreamInfo: StreamInfo{
					WorkerId:    "worker-1",
					ServiceName: "service-1",
				},
				Values: [][]interface{}{
					{
						timeString,
						"fizzbuzz-1",
					},
					{
						timeString,
						"fizzbuzz-1",
					},
				},
			},
		},
	}

	// 将数据编码为 JSON 字符串
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
		return
	}
	fmt.Printf("%+v\n", string(jsonData))

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 解析响应
	fmt.Println("Response Status:", resp.Status)
}
