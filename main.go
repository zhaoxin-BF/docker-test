package main

import (
	"fmt"
	"github.com/zhaoxin-BF/docker-test/container"
)

func main() {
	//common.GetTimezone()
	//common.ListTimeZones()
	//common.GetOSTimezone()
	//common.GetTZ()

	//container.GetcContainer()

	fmt.Println("-----------------------------------: print logs")
	//container.GetContainerLogs()
	//container.GetStreamLogs()
	//container.ContainerStats()
	//container.StreamRead()
	container.ContainerTop()

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

	// loki log
	//grafana_loki.GetLog()

	// Get Location
	//local_time.GetLocation()
	//local_time.GetLocationHttp()
}
