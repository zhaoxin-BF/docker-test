package main

import (
	"fmt"
	"net/http"
	"time"
)

func Do() {
	client := http.Client{
		Timeout: 5 * time.Second, // 设置超时时间为 5 秒
	}

	request, err := http.NewRequest("GET", "http://127.0.0.1:6680/fs", nil)

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("请求发生错误:", err)
		return
	}

	defer resp.Body.Close()

	// 处理响应...

	fmt.Printf("%+v\n", resp.Status)
	fmt.Printf("%+v\n", resp.StatusCode)

	fmt.Printf("%+v\n", resp)
}
