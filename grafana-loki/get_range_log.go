package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type rangeData struct {
	Status string      `json:"status"`
	Data   rangeResult `json:"data"`
}

type rangeResult struct {
	Result     []rangeValUnit `json:"result"`
	ResultType string         `json:"resultType"`
}

type rangeValUnit struct {
	Stream map[string]string `json:"stream"`
	Values [][]any           `json:"values"`
}

func LokiGetLogRange() {
	// Loki 地址
	lokiURL := "http://192.168.31.154:3100/loki/api/v1/query_range"

	fmt.Println(time.Now().Unix() - 3600*45)
	fmt.Println(time.Now().Unix())

	// 设置查询参数
	params := fmt.Sprintf("query=%s&start=%d&end=%d&step=%d",
		"{container_name=\"boreas-nginx\"}",
		//"{worker_id=\"worker-1\"}",
		time.Now().Unix()-3600*45,
		time.Now().Unix(),
		6000)

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

	res := &rangeData{}

	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("反序列化错误:", err)
		return
	}

	//fmt.Printf("%+v", res.Data.Result)

	for _, v := range res.Data.Result {
		fmt.Println(v.Stream)
		for _, v2 := range v.Values {
			fmt.Println(v2)
			fmt.Println(v2[0])
			fmt.Println(v2[1])
		}
	}

	// 打印日志
	//fmt.Println(string(body))
}

func LokiGetLog() {
	// Loki 地址
	lokiURL := "http://192.168.31.154:3100/loki/api/v1/query"

	// 设置查询参数
	params := fmt.Sprintf("query=%s&time=1715769569",
		"{container_name=\"boreas-nginx\"}")

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

	res := &rangeData{}

	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("反序列化错误:", err)
		return
	}

	fmt.Printf("%+v", res.Data.Result)

	for _, v := range res.Data.Result {
		fmt.Println(v.Stream)
		for _, v2 := range v.Values {
			//fmt.Println(v2)
			fmt.Println(v2[0])
			//fmt.Println(v2[1])
		}
	}

	// 打印日志
	//fmt.Println(string(body))
}
