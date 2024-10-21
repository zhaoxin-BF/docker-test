package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Req struct {
	Code    int
	UseTime int64
}

func main() {
	oneReq := 60
	times := 10
	reqs := make(chan Req, oneReq*times)

	wg := &sync.WaitGroup{}
	for j := 0; j < times; j++ {
		wg.Add(1)
		go func() {
			for i := 0; i < oneReq; i++ {
				reqZhuiyi(reqs, wg)
			}
		}()
	}

	wg.Wait()

	close(reqs)

	var maxUseTime, successTotal, failedTotal, minUseTime, successTotalTime int64

	for v := range reqs {
		if v.Code == http.StatusOK {
			successTotal++
			successTotalTime += v.UseTime
			if minUseTime == 0 {
				minUseTime = v.UseTime
			}
			if v.UseTime < minUseTime {
				minUseTime = v.UseTime
			}
			if v.UseTime > maxUseTime {
				maxUseTime = v.UseTime
			}
		} else {
			failedTotal++
		}
	}

	fmt.Println("成功的次数: ", successTotal)
	fmt.Println("失败的次数: ", failedTotal)
	fmt.Println("最大耗时: ", maxUseTime)
	fmt.Println("最小耗时: ", minUseTime)
	fmt.Println("平均耗时: ", successTotalTime/successTotal)

}

func reqZhuiyi(code chan Req, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	body := &bytes.Buffer{}
	_, err := body.Write([]byte(`{"model": "Qwen2-1.5B-msft-v2-simpletod-v7-endofcall", "messages": [{"role": "user", "content": "你是谁"}], "stream": true}`))
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPost, "https://everai.expvent.com/r/apps/zhuiyi-bowen-serve/v1/chat/completions", body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer everai_DmCZ8VZhwc6Dx4kX7uHUiO")

	httpClient := http.DefaultClient
	startTime := time.Now()
	//fmt.Println("开始请求")
	resp, err := httpClient.Do(req)
	//fmt.Println("请求结束")
	endTime := time.Now().Sub(startTime).Milliseconds()
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//respBody, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("resp body read err: ", err)
	//} else {
	//	fmt.Println("resp body: ", string(respBody))
	//}
	//fmt.Printf("resp code:%v; resp header: %#v, ", resp.StatusCode, resp.Header)
	fmt.Printf("UseTime: %d ms\n", endTime)
	code <- Req{Code: resp.StatusCode, UseTime: endTime}
}
