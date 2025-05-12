package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// 发送单个请求的函数
func sendRequest(url string, semaphore chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	// 获取信号量，限制并发数
	semaphore <- struct{}{}
	defer func() { <-semaphore }()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("请求出错: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应体出错: %v\n", err)
		return
	}

	fmt.Printf("响应内容: %s\n", body)
}

func main() {
	url := "https://everai.expvent.com/r/apps/beifeng-001/hello"
	requestCount := 10000 // 发送请求的数量
	concurrency := 100    // 并发数

	var wg sync.WaitGroup
	// 使用通道实现信号量
	semaphore := make(chan struct{}, concurrency)

	for i := 0; i < requestCount; i++ {
		wg.Add(1)
		go sendRequest(url, semaphore, &wg)
	}

	wg.Wait()
}
