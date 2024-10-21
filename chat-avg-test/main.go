package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ResponseTime struct {
	time int64
}

func sendRequest() (int64, error) {
	url := "http://36.213.13.203:6688/v1/chat/completions"
	//url := "https://everai.expvent.com/r/apps/zhuiyi-bowen-serve/v1/chat/completions"
	headers := http.Header{
		"Content-Type":       []string{"application/json"},
		"x-everai-worker-id": []string{"b4WH6nKjZoCTvQbDpVAaop:6701"},
		"authorization":      []string{"Bearer everai_DmCZ8VZhwc6Dx4kX7uHUiO"},
	}
	data := []byte(`{"model":"Qwen2-1.5B-msft-v2-simpletod-v7-endofcall","messages":[{"role":"user","content":"你是谁"}],"stream":true}`)

	startTime := time.Now()
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return 0, err
	}
	req.Header = headers
	req.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return 0, err
	}

	endTime := time.Now()
	return endTime.Sub(startTime).Milliseconds(), nil
}

func main() {
	var times []ResponseTime
	for i := 0; i < 20; i++ {
		elapsedTime, err := sendRequest()
		if err != nil {
			log.Printf("Request %d error: %v", i+1, err)
			continue
		}
		fmt.Println(elapsedTime, "ms")
		times = append(times, ResponseTime{elapsedTime})
	}

	var totalTime int64
	minTime := times[0].time
	maxTime := times[0].time

	for _, t := range times {
		if t.time < minTime {
			minTime = t.time
		}
		if t.time > maxTime {
			maxTime = t.time
		}
		totalTime += t.time
	}

	if len(times) > 0 {
		avgTime := totalTime / int64(len(times))
		log.Printf("Minimum time: %dms", minTime)
		log.Printf("Maximum time: %dms", maxTime)
		log.Printf("Average time: %dms", avgTime)
	} else {
		log.Println("No successful requests to calculate statistics.")
	}
}
