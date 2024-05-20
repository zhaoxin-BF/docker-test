package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type RequestBody struct {
	Prompt string `json:"prompt"`
}

func main() {
	// 设置总请求数量
	totalRequests := 1000

	// 设置并发数量
	concurrency := 10

	// 创建 WaitGroup 用于等待所有请求完成
	var wg sync.WaitGroup
	wg.Add(totalRequests)

	// 创建一个 channel 用于限制并发数量
	semaphore := make(chan struct{}, concurrency)

	// 发送并发请求
	startTime := time.Now().Unix()
	for i := 0; i < totalRequests; i++ {
		semaphore <- struct{}{}
		fmt.Printf("------------------------------- 第: %d 次请求。\n", i)
		go func() {
			// 获取信号量
			defer func() {
				// 释放信号量
				<-semaphore
				wg.Done()
			}()
			sendRequest()
		}()
	}
	// 等待所有请求完成
	wg.Wait()
	endTime := time.Now().Unix()
	fmt.Printf("======== 总耗时：%+vs ========", endTime-startTime)
}

func sendRequest() {
	// 请求体
	requestBody := RequestBody{
		Prompt: "who are you",
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Printf("Error marshaling request body: %v\n", err)
		return
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", "https://everai.expvent.com.cn:1111/api/apps/v1/routes/test-llama2-7b-chat-modal-2/chat", bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Printf("Error creating HTTP request: %v\n", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer everai_637wE9obZtmGLyqIJp0lok")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending HTTP request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 打印响应状态码
	fmt.Printf("Response status: %d\n", resp.StatusCode)
}
