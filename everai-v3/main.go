package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// 文件路径
	filePath := "/Users/zhaoxin/Workspace/github.com/zhaoxin-BF/docker-test/everai-v3/helm-v3.4.2-linux-amd64.tar.gz"

	// 上传 URL
	uploadURL := "https://everai.expvent.com/r/apps/everai-test-go-003/upload"

	// 记录开始时间
	startTime := time.Now()

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 创建 POST 请求
	req, err := http.NewRequest("POST", uploadURL, file)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "multipart/form-data")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 记录结束时间并计算耗时
	duration := time.Since(startTime)

	// 输出结果
	if resp.StatusCode == http.StatusOK {
		fmt.Printf("File %s uploaded successfully in %v\n", filepath.Base(filePath), duration)
	} else {
		fmt.Println("Error uploading file:", resp.Status)
	}
}
