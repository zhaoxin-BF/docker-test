package grafana_loki

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GetLog() {
	// Loki 地址
	lokiURL := "http://192.168.31.154:3100/loki/api/v1/query_range"

	// 设置查询参数
	params := fmt.Sprintf("query=%s&start=%d&end=%d&step=%d",
		"{container_name=\"nginx\"}",
		time.Now().Unix()-3600*24,
		time.Now().Unix(),
		15*60)

	// 创建请求
	req, err := http.NewRequestWithContext(context.Background(), "GET", lokiURL+"?"+params, nil)
	if err != nil {
		panic(err)
	}

	// 发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var unJson map[string]interface{}

	err = json.Unmarshal(body, &unJson)
	if err != nil {
		fmt.Println("反序列化错误:", err)
		return
	}

	fmt.Printf("%+v\n", unJson)
	// 打印日志
	//fmt.Println(string(body))
}
