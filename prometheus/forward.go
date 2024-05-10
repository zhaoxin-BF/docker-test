package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Forward() {
	// 发起 GET 请求获取指标数据
	resp, err := http.Get("http://192.168.31.154:9100/metrics")
	if err != nil {
		fmt.Println("GET请求失败:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	// 去除重复的指标数据
	uniqueBody := removeDuplicates(body)

	// 发起 POST 请求将指标数据发送到Pushgateway
	pushgatewayURL := "http://192.168.31.154:9091/metrics/job/agent-server/deviceId/id-xxxxxx"
	resp, err = http.Post(pushgatewayURL, "text/plain", bytes.NewReader(uniqueBody))
	if err != nil {
		fmt.Println("POST请求失败:", err)
		return
	}
	body, err = ioutil.ReadAll(resp.Body)

	fmt.Printf("%+v", string(body))
	fmt.Println("指标数据已成功推送到Pushgateway")
}

// Duplicate indicator data is removed
func removeDuplicates(data []byte) []byte {
	lines := strings.Split(string(data), "\n")
	uniqueLines := make(map[string]struct{})

	var uniqueData []byte
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" {
			if _, ok := uniqueLines[trimmedLine]; !ok {
				uniqueLines[trimmedLine] = struct{}{}
				uniqueData = append(uniqueData, trimmedLine...)
				uniqueData = append(uniqueData, '\n')
			}
		}
	}
	return uniqueData
}
