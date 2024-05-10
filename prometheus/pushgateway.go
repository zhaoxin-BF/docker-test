package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func main01() {
	PushGateway()
}

func PushGateway() {
	completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_last_completion_timestamp_seconds",
		Help: "The timestamp of the last successful completion of a DB backup.",
	})
	completionTime.SetToCurrentTime()
	// completionTime.Set(200)  // set可以设置任意值（float64）

	pusher := push.New("http://192.168.31.154:9091", "db_backup"). // push.New("pushgateway地址", "job名称")
									Collector(completionTime).                                   // Collector(completionTime) 给指标赋值
									Grouping("db", "customers").Grouping("instance", "1.1.1.1"). // 给指标添加标签，可以添加多个
									BasicAuth("", "")

	if err := pusher.Push(); err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}
}
