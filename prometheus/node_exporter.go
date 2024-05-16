package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		// 设置 Content-Type 为 text/plain
		w.Header().Set("Content-Type", "text/plain")

		// 发送监控参数 test-count = 6
		fmt.Fprintf(w, "# HELP test_count This is a test count metric\n")
		fmt.Fprintf(w, "# TYPE test_count gauge\n")
		fmt.Fprintf(w, "test_count{instance=\"localhost\"} %s\n", "6")
	})

	// 启动 HTTP 服务器并监听端口 9600
	fmt.Println("Starting node_exporter demo on port 9600...")
	log.Fatal(http.ListenAndServe(":9600", nil))
}

type Metric struct {
	name   string
	Lables map[string]string
	Mumber int64
}

func test(m Metric) {
	//metric := "everai_container_spec_cpu_percentage{containerName=\"worker-11\", workerId=\"worker-11\", appId=\"app-11\"} 57\n"
}

func formatMetric(m Metric) string {
	labels := ""
	for k, v := range m.Lables {
		labels += fmt.Sprintf("%s=\"%s\", ", k, v)
	}
	labels = strings.TrimSuffix(labels, ", ")

	return fmt.Sprintf("%s{%s} %d\n", m.name, labels, m.Mumber)
}
