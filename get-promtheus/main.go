package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// 定义 Prometheus API 端点
	apiURL := "http://192.168.31.43:31169/api/v1/query"

	// 定义要查询的指标
	metrics := []string{"everai_container_spec_cpu_percentage", "everai_container_spec_memory_percentage"}

	// 构建查询参数
	values := url.Values{}
	for _, metric := range metrics {
		values.Add("query", metric)
	}

	// 发送 HTTP GET 请求
	resp, err := http.Get(apiURL + "?" + values.Encode())
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	// 解析 JSON 响应
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("解析 JSON 失败:", err)
		return
	}

	// 提取指标数值并打印
	results, ok := data["data"].(map[string]interface{})["result"].([]interface{})
	if ok {
		for _, metric := range metrics {
			for _, result := range results {
				resultMap, ok := result.(map[string]interface{})
				if ok {
					metricName, ok := resultMap["metric"].(map[string]interface{})["__name__"].(string)
					if ok && metricName == metric {
						value, ok := resultMap["value"].([]interface{})[1].(string)
						if ok {
							fmt.Printf("%s: %s\n", metric, value)
						}
					}
				}
			}
		}
	}
}
