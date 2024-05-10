package local_time

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"net/http"
)

func GetLocation() {
	client := resty.New()
	resp, err := client.R().Get("http://ip-api.com/json")
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}

	location := resp.String()
	fmt.Println("机器物理定位信息:", location)
}

func GetLocationHttp() {
	resp, err := http.Get("http://ip-api.com/json")
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	location := string(body)
	fmt.Println("机器物理定位信息:", location)
}
