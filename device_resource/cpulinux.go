package device_resource

import (
	"context"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
)

type CpuInfo struct {
	Model        string // CPU model, for example: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
	Cores        uint32 // CPU core, for example: CPU(s): 12
	Threads      uint32 // CPU Thread(s) per core，for example：Thread(s) per core:  2
	Architecture string // CPU arch, for example: x86_64/arm64/~
	CpuNum       uint32 // Number of cpus, In most cases, there will only be one physical cpu
}

type CpuCollector struct {
	CPU CpuInfo
}

func NewCPUCollector() *CpuCollector {
	return &CpuCollector{}
}

func (c *CpuCollector) GetCPUInfo() (*CpuInfo, error) {
	var cpuInfo CpuInfo
	info, err := cpu.InfoWithContext(context.Background())
	if err != nil {
	}

	// test printf
	fmt.Printf("%+v", info)
	var cpuPhysicalMap = make(map[string][]cpu.InfoStat)

	for _, val := range info {
		key := val.PhysicalID
		if val.PhysicalID == "" {
			key = "null"
		}
		cpuPhysicalMap[key] = append(cpuPhysicalMap[key], val)
		cpuInfo.Cores += uint32(val.Cores)
		cpuInfo.Model = val.ModelName
	}
	cpuInfo.CpuNum = uint32(len(cpuPhysicalMap))

	str, err := host.HostID()
	fmt.Println("_______________hostId: ", str)

	hostInfo, err := host.Info()
	fmt.Println("%+v", hostInfo)

	//for _, val := range info {
	//	fmt.Println("CPU 信息:")
	//	fmt.Printf("  CPU 数量: %d\n", val.CPU)
	//	fmt.Printf("  厂商 ID: %s\n", val.VendorID)
	//	fmt.Printf("  系列: %s\n", val.Family)
	//	fmt.Printf("  型号: %s\n", val.Model)
	//	fmt.Printf("  Stepping: %d\n", val.Stepping)
	//	fmt.Printf("  物理 ID: %s\n", val.PhysicalID)
	//	fmt.Printf("  核心 ID: %s\n", val.CoreID)
	//	fmt.Printf("  核心数量: %d\n", val.Cores)
	//	fmt.Printf("  型号名称: %s\n", val.ModelName)
	//	fmt.Printf("  主频: %.2f MHz\n", val.Mhz)
	//	fmt.Printf("  缓存大小: %d KB\n", val.CacheSize)
	//	fmt.Printf("  Flags: %v\n", val.Flags)
	//	fmt.Printf("  微码: %s\n", val.Microcode)
	//}
	return &cpuInfo, nil
}
