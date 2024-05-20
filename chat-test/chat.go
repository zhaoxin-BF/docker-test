package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type RequestBody struct {
	Prompt string `json:"prompt"`
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExample:\n")
		fmt.Fprintf(os.Stderr, "  %s -client-count 100 -requests-per-client 10\n", os.Args[0])
	}
}

func main() {
	var clientCount, requestsPerClient int
	flag.IntVar(&clientCount, "client-count", 0, "The number of concurrent customers")
	flag.IntVar(&requestsPerClient, "requests-per-client", 0, "The number of individual customer requests")
	flag.Parse()

	if clientCount <= 0 || requestsPerClient <= 0 {
		flag.Usage()
		os.Exit(1)
	}
	var wg sync.WaitGroup

	startTime := time.Now().Unix()
	for i := 0; i < clientCount; i++ {
		wg.Add(1)
		go func(c int) {
			defer wg.Done()
			Customer(c, requestsPerClient)
		}(i)
	}

	wg.Wait()

	endTime := time.Now().Unix()
	fmt.Printf("\n\n\n\n")
	fmt.Printf("======== Client Count: %d, Per Client Request Count: %d, Total Request: %d ======\n", clientCount, requestsPerClient, clientCount*requestsPerClient)
	fmt.Printf("======== Total Time Spent：%+vs ========\n", endTime-startTime)
}

func Customer(c int, totalRequests int) {
	// send request
	startTime := time.Now().Unix()
	for i := 0; i < totalRequests; i++ {
		sendRequest(c, i)
	}
	// wait
	endTime := time.Now().Unix()
	fmt.Printf("Customer %d ======== Total time spent：%+vs ========\n", c, endTime-startTime)
}

func sendRequest(c int, times int) {
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
	startTime := time.Now().Unix()
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

	endTime := time.Now().Unix()
	// 打印响应状态码
	fmt.Printf("Customer %d, times %d, Response status: %d, per_request_time_spent %ds\n", c, times, resp.StatusCode, endTime-startTime)
}
