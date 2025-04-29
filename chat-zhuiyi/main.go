//package main
//
//import (
//	"bytes"
//	"flag"
//	"fmt"
//	"io/ioutil"
//	"log"
//	"net/http"
//	"sync"
//)
//
//func sendRequest(url string, headers http.Header, data []byte) (string, int) {
//	req, err := http.NewRequest("POST", url, nil)
//	if err != nil {
//		log.Printf("Error creating request: %v", err)
//		return "", 0
//	}
//	req.Header = headers
//	req.Body = ioutil.NopCloser(bytes.NewBuffer(data))
//
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Printf("Error sending request: %v", err)
//		return "", 0
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Printf("Error reading response body: %v", err)
//		return "", resp.StatusCode
//	}
//
//	return string(body), resp.StatusCode
//}
//
//func runRequests(n, c int) {
//	url := "https://everai.expvent.com/r/apps/zhuiyi-bowen-serve/v1/chat/completions"
//	headers := http.Header{
//		"Content-Type":  []string{"application/json"},
//		"authorization": []string{"Bearer everai_DmCZ8VZhwc6Dx4kX7uHUiO"},
//	}
//	data := []byte(`{"model":"Qwen2-1.5B-msft-v2-simpletod-v7-endofcall","messages":[{"role":"user","content":"你是谁"}],"stream":true}`)
//
//	var wg sync.WaitGroup
//	ch := make(chan struct{}, c)
//	for i := 1; i <= n; i++ {
//		wg.Add(1)
//		ch <- struct{}{}
//		go func(index int) {
//			defer wg.Done()
//			response, statusCode := sendRequest(url, headers, data)
//			fmt.Printf("Request %d: Status Code: %d, Response: %s\n", index, statusCode, response)
//			<-ch
//		}(i)
//	}
//	wg.Wait()
//}
//
//func main() {
//	var n, c int
//	flag.IntVar(&n, "n", 0, "Total number of requests")
//	flag.IntVar(&c, "c", 0, "Concurrency level")
//	flag.Parse()
//
//	if n <= 0 || c <= 0 {
//		fmt.Println("Please provide valid values for n and c.")
//		return
//	}
//
//	fmt.Println("Starting requests...")
//	runRequests(n, c)
//}

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"sync"
	"time"
)

func sendRequest(client *http.Client, url, method string, headers http.Header, data []byte, clientIndex, requestIndex int) int {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Printf("Client %d, Request %d: Error creating request: %v", clientIndex, requestIndex, err)
		return 0
	}
	req.Header = headers
	req.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Client %d, Request %d: Error sending request: %v", clientIndex, requestIndex, err)
		return 0
	}
	defer resp.Body.Close()
	return resp.StatusCode
}

func saveResponseToFile(clientIndex, requestIndex int, response string) {
	folderName := fmt.Sprintf("client_%d", clientIndex)
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		log.Printf("Error creating folder for client %d: %v", clientIndex, err)
		return
	}

	fileName := fmt.Sprintf("%s/response_%d-%d.txt", folderName, clientIndex, requestIndex)
	err = ioutil.WriteFile(fileName, []byte(response), os.ModePerm)
	if err != nil {
		log.Printf("Error writing to file for client %d, request %d: %v", clientIndex, requestIndex, err)
	}
}

type Req struct {
	Code    int
	UseTime int64
}

func runRequests(n, c int, url, token, body, method string, sleepMs int) {
	//url := "https://everai.expvent.com/r/apps/everai-bowen-serve/v1/chat/completions"
	//url = "https://everai.expvent.com/r/apps/zhuiyi-bowen-serve/v1/chat/completions" // 追一
	//token = "everai_DmCZ8VZhwc6Dx4kX7uHUiO"
	//body = "{\"model\":\"Qwen2-1.5B-msft-v2-simpletod-v7-endofcall\",\"messages\":[{\"role\":\"user\",\"content\":\"你是谁\"}],\"stream\":true}"

	headers := http.Header{
		"Content-Type": []string{"application/json"},
		//"authorization": []string{"Bearer everai_3pdgNzSkA9LColW5dOzgb2"},
		"authorization": []string{fmt.Sprintf("Bearer %s", token)}, // 追一
	}
	data := []byte(body)

	// 发送请求
	transport := &http.Transport{
		MaxIdleConns:        5000,
		MaxIdleConnsPerHost: 5000,
		IdleConnTimeout:     10 * time.Second,
	}

	client := &http.Client{
		//Timeout:   300 * time.Second,
		Transport: transport,
	}

	reqs := make(chan Req, c*n)

	TstartTime := time.Now().Unix()
	var wg sync.WaitGroup
	for clientIndex := 1; clientIndex <= c; clientIndex++ {
		//for requestIndex := 1; requestIndex <= n/c; requestIndex++ {
		wg.Add(1)
		go func(code chan Req, c, requests int) {
			defer wg.Done()
			for i := 1; i <= requests; i++ {
				// 后进行休眠
				time.Sleep(time.Duration(sleepMs) * time.Millisecond)
				startTime := time.Now()
				statusCode := sendRequest(client, url, method, headers, data, c, i)
				endTime := time.Now()
				elapsedTime := endTime.Sub(startTime).Milliseconds()
				fmt.Printf("Client %d, Request %d: Status Code: %d, Time: %dms\n", c, i, statusCode, elapsedTime)
				//go saveResponseToFile(client, i, response)
				code <- Req{Code: statusCode, UseTime: elapsedTime}
			}
		}(reqs, clientIndex, n)
	}
	wg.Wait()
	TendTime := time.Now().Unix()

	requestPerTime := int64(c*n) / (TendTime - TstartTime)

	close(reqs)
	var maxUseTime, successTotal, failedTotal, minUseTime, successTotalTime int64
	var successTimes []int64

	// 遍历每个请求
	for v := range reqs {
		if v.Code == http.StatusOK {
			// 统计成功请求的总数
			successTotal++
			// 累加成功请求的使用时间
			successTotalTime += v.UseTime
			// 将成功请求的使用时间添加到切片中
			successTimes = append(successTimes, v.UseTime)

			// 记录最小和最大耗时
			if minUseTime == 0 || v.UseTime < minUseTime {
				minUseTime = v.UseTime // 更新最小耗时
			}
			if v.UseTime > maxUseTime {
				maxUseTime = v.UseTime // 更新最大耗时
			}
		} else {
			// 统计失败请求的总数
			failedTotal++
		}
	}

	// 输出成功和失败的请求次数
	fmt.Printf("\n\n\n请求总次数: %d \n", c*n)
	fmt.Printf("并发数：%d \n", c)
	fmt.Printf("成功的次数: %d \n", successTotal)
	fmt.Printf("失败的次数: %d \n", failedTotal)
	fmt.Printf("总耗时(sec): %d \n", TendTime-TstartTime)
	fmt.Printf("每秒平均请求数(#/sec): %d \n", requestPerTime)
	fmt.Printf("最大耗时(ms): %d \n", maxUseTime)
	fmt.Printf("最小耗时(ms): %d \n", minUseTime)

	// 计算并输出平均耗时
	if successTotal > 0 {
		averageTime := successTotalTime / successTotal // 计算平均耗时
		fmt.Printf("平均耗时(ms): %d \n", averageTime)

		// 计算并输出百分位数数据
		sort.Slice(successTimes, func(i, j int) bool {
			return successTimes[i] < successTimes[j] // 按照使用时间排序
		})

		// 定义需要计算的百分位数
		fmt.Println("\nPercentage of the requests served within a certain time (ms)")
		percentiles := []int{50, 66, 75, 80, 90, 95, 98, 99, 100}
		for _, p := range percentiles {
			index := (p * int(successTotal) / 100) - 1 // 计算百分位数的索引
			if index < 0 {
				index = 0 // 确保索引不小于0
			}
			// 输出百分位数和对应的使用时间
			fmt.Printf("%3d%%    %d\n", p, successTimes[index])
		}
	} else {
		// 如果没有成功的请求，输出提示信息
		fmt.Println("没有成功的请求，无法计算平均耗时和百分比数据")
	}
}

func init() {
	// 设置flag的使用方式
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExample:\n")
		fmt.Fprintf(os.Stderr, "  %s -client-count 2 -requests-per-client 10 -token 'everai_637wE9obZtmGLyqIJp0lok' -url http://example.com -method POST -body 'hello'\n", os.Args[0])
	}
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

}

func main() {
	var clientCount, requestsPerClient int
	var token, url, method, body string
	var sleepMs int
	flag.IntVar(&clientCount, "client-count", 0, "The number of concurrent customers")
	flag.IntVar(&requestsPerClient, "requests-per-client", 0, "The number of individual customer requests")
	flag.StringVar(&url, "url", "", "request url")
	flag.StringVar(&method, "method", "POST", "request method")
	flag.StringVar(&body, "body", "", "request body")
	flag.StringVar(&token, "token", "", "token for authentication")
	flag.IntVar(&sleepMs, "sleep-ms", 0, "sleep time before per request (ms)")
	flag.Parse()

	log.Printf("Executing with parameters: clientCount=%d, requestsPerClient=%d, url=%s, method=%s, body=%s", clientCount, requestsPerClient, url, method, body)

	if clientCount == 0 {
		log.Fatalf("client-count is required")
	}
	if requestsPerClient == 0 {
		log.Fatalf("requests-per-client is required")
	}
	if url == "" {
		log.Fatalf("url is required")
	}
	if method == "" {
		log.Fatalf("method is required")
	}
	if body == "" {
		log.Fatalf("body is required")
	}

	fmt.Println("Starting requests...")
	runRequests(requestsPerClient, clientCount, url, token, body, method, sleepMs)
}
