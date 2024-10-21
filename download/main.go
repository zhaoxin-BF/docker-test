package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func downloadFile(url string) time.Duration {
	startTime := time.Now()
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error downloading file: %v\n", err)
		return 0
	}
	defer response.Body.Close()

	endTime := time.Now()
	return endTime.Sub(startTime)
}

func download() {
	// 解析命令行参数
	downloadUrl := "http://140.143.153.97:8080/download/response.txt"
	//downloadUrl := "http://36.213.13.233:7008/download/response.txt"

	// 并发下载文件
	downloadDurations := []time.Duration{}
	var wg sync.WaitGroup
	ch := make(chan time.Duration, 1)
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			duration := downloadFile(downloadUrl)
			ch <- duration
		}()
	}
	wg.Wait()
	close(ch)
	for duration := range ch {
		downloadDurations = append(downloadDurations, duration)
	}

	// 计算平均下载时间
	var total time.Duration
	for _, duration := range downloadDurations {
		total += duration
	}
	avgDownloadTime := total / time.Duration(1)

	// 输出结果
	fmt.Printf("Average download time per file: %.2f seconds\n", avgDownloadTime.Seconds())
}

func main() {
	download()
}
