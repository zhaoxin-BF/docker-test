package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type RequestBody struct {
	Prompt string `json:"prompt"`
}

var fileSuffix = map[string]string{
	"chats":  ".txt",
	"images": ".png",
}

func init() {
	// 设置flag的使用方式
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExample:\n")
		fmt.Fprintf(os.Stderr, "  %s -client-count 100 -requests-per-client 10 -token 'Bearer everai_637wE9obZtmGLyqIJp0lok' -class chats -url http://example.com -method POST -prompt 'hello'\n", os.Args[0])
	}
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

}

func main() {
	var clientCount, requestsPerClient int
	var token, url, method, prompt, class string
	var save bool
	flag.IntVar(&clientCount, "client-count", 0, "The number of concurrent customers")
	flag.IntVar(&requestsPerClient, "requests-per-client", 0, "The number of individual customer requests")
	flag.StringVar(&class, "class", "chats", "The request class chats/images")
	flag.StringVar(&url, "url", "", "request url")
	flag.StringVar(&method, "method", "POST", "request method")
	flag.StringVar(&prompt, "prompt", "", "prompt")
	flag.StringVar(&token, "token", "", "token")
	flag.BoolVar(&save, "save", true, "whether save response")
	flag.Parse()

	log.Printf("Executing with parameters: clientCount=%d, requestsPerClient=%d, url=%s, method=%s, prompt=%s", clientCount, requestsPerClient, url, method, prompt)

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
	//if token == "" {
	//	log.Fatalf("token is required")
	//}
	//if prompt == "" {
	//	log.Fatalf("prompt is required")
	//}
	if class == "chats" || class == "images" {
		// nothing
	} else {
		log.Fatalf("class is not allowed")
		return
	}

	var wg sync.WaitGroup

	startTime := time.Now().Unix()
	for i := 0; i < clientCount; i++ {
		wg.Add(1)
		go func(c int) {
			defer wg.Done()
			Customer(token, method, url, prompt, class, c, requestsPerClient, save)
		}(i)
	}

	wg.Wait()

	endTime := time.Now().Unix()
	fmt.Printf("\n\n\n\n")
	fmt.Printf("======== Client Count: %d, Per Client Request Count: %d, Total Request: %d ======\n", clientCount, requestsPerClient, clientCount*requestsPerClient)
	fmt.Printf("======== Total Time Spent：%+vs ========\n", endTime-startTime)
}

func Customer(token, method, url, prompt, class string, c int, totalRequests int, save bool) {
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
	// send request
	//startTime := time.Now().Unix()
	for i := 0; i < totalRequests; i++ {
		sendRequest(client, token, method, url, prompt, class, c, i, save)
	}
	// wait
	//endTime := time.Now().Unix()
	//fmt.Printf("Customer %d ======== Total time spent：%+vs ========\n", c, endTime-startTime)
}

func sendRequest(client *http.Client, token, method, url, prompt, class string, c int, times int, save bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Customer %d ======== Recover Error: %+v\n", c, r)
		}
	}()
	// 请求体
	requestBody := RequestBody{
		Prompt: prompt,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Printf("Error marshaling request body: %v\n", err)
		return
	}

	// 创建 HTTP 请求
	startTime := time.Now().Unix()
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Printf("Error creating HTTP request: %v\n", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending HTTP request: %v\n", err)
		return
	}
	endTime := time.Now().Unix()
	fmt.Printf("Customer %d, times %d, per_request_time_spent %ds\n", c, times, endTime-startTime)
	resp.Body.Close()
	// 进一步确保连接被关闭
	if resp.Body != nil {
		resp.Body.Close()
	}
	//if save {
	//	go func(resp *http.Response) {
	//		defer func() {
	//			if r := recover(); r != nil {
	//				fmt.Printf("Recovered in customer %d: %v\n", c, resp)
	//			}
	//			if err := resp.Body.Close(); err != nil {
	//				fmt.Printf("Error closing response body: %v\n", err)
	//			}
	//		}()
	//		if resp.StatusCode != http.StatusOK {
	//			fmt.Printf("Customer %d, times %d, Response status failed, status %d\n", c, times, resp.StatusCode)
	//			body, err := ioutil.ReadAll(resp.Body)
	//			if err != nil {
	//				fmt.Printf("Error reading response body: %v\n", err)
	//				return
	//			}
	//			fmt.Printf("Response body: %s\n", string(body))
	//			return
	//		}
	//
	//		// 创建用于保存响应 body 的目录
	//		outputDir := "Customer-" + strconv.Itoa(c)
	//		err = os.Mkdir(outputDir, 0755)
	//		if err != nil && !os.IsExist(err) {
	//			log.Printf("Failed to create output directory: %v", err)
	//		}
	//
	//		// 将响应 body 保存到文件
	//		fileName := fmt.Sprintf("response_%d%s", times, fileSuffix[class])
	//		filePath := filepath.Join(outputDir, fileName)
	//		file, err := os.Create(filePath)
	//		if err != nil {
	//			log.Printf("Failed to create file: %v", err)
	//		}
	//		_, err = io.Copy(file, resp.Body)
	//		if err != nil {
	//			log.Printf("Failed to write response body to file: %v", err)
	//			return
	//		}
	//		err = file.Close()
	//		if err != nil {
	//			log.Printf("Failed to close file: %v", err)
	//			return
	//		}
	//	}(resp)
	//}
}
