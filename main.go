package main

import (
	"fmt"
	"time"
)

func main() {
	//common.GetTimezone()
	//common.ListTimeZones()
	//common.GetOSTimezone()
	//common.GetTZ()

	//container.GetcContainer()

	currentTime := time.Now().UTC()
	twoHoursAgo := currentTime.Add(-2 * time.Hour)
	timeString := twoHoursAgo.Format("2006-01-02T15:04:05.0Z")
	fmt.Println(timeString)

	fmt.Println("-----------------------------------: print logs")
	//container.GetContainerLogs()
	//container.GetStreamLogs()
	//container.ContainerStats()
	//container.StreamRead()
	//container.ContainerTop()

	//device_resource.GetDeviceResource()

	// get cpu info
	//device_resource_new.GetCPUInfo()

	//fmt.Println("hello world")

	//newDir.Do()

	// config
	//config.GetConfig()

	// prometheus
	//prometheus.Forward()

	// docker log
	//docker_log.GetLog()

	//loki log
	fmt.Println("-----------------------------------: loki get logs")
	//grafana_loki.LokiTailLog()
	//grafana_loki.LokiGetLog()
	//grafana_loki.LokiGetLogRange()
	//grafana_loki.PushLogIntoLoki()

	// Get Location
	//local_time.GetLocation()
	//local_time.GetLocationHttp()
}
