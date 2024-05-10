package device_resource

import (
	"fmt"
)

func GetDeviceResource() {
	////get filesystem
	//fmt.Printf("---------------------------------------------get filesystem:\n")
	//filesystemCollector := NewFilesystemCollector()
	//filesystemInfos, err := filesystemCollector.GetFilesystemInfo()
	//if err != nil {
	//}
	//for _, val := range filesystemInfos {
	//	fmt.Println(val.Device, "-----", val.Type, "------", val.Uuid)
	//}
	//
	//get memory
	//fmt.Printf("---------------------------------------------get memory:\n")
	//memoryCollector := NewMemoryCollector()
	//memoryInfo, _ := memoryCollector.GetMemoryInfo()
	//fmt.Println(memoryInfo.Total)
	//fmt.Println(memoryInfo.Used)
	//fmt.Println(memoryInfo.Free)
	//fmt.Println(memoryInfo)
	//
	////get cpu
	//fmt.Printf("---------------------------------------------get cpu:\n")
	//cpu := NewCPUCollector()
	//cpu.GetCPUInfo()
	////fmt.Println(cpuinfo)

	// get GPU
	fmt.Printf("---------------------------------------------get GPU:\n")
	gpu := NewGPUCollector()
	gpu.GetGPUInfo()
}
