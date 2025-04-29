package device_resource_new

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
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
	cmd := exec.CommandContext(context.Background(), "lscpu")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("执行 lscpu 命令失败:", err)
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	cpuInfo := &CpuInfo{}
	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "Model name":
			cpuInfo.Model = value
		case "CPU(s)":
			cores, err := strconv.ParseUint(value, 10, 32)
			if err != nil {
				fmt.Printf("解析 CPU(s) 值失败: %v\n", err)
				return nil, err
			}
			cpuInfo.Cores = uint32(cores)
		case "Thread(s) per core":
			threads, err := strconv.ParseUint(value, 10, 32)
			if err != nil {
				fmt.Printf("解析 Thread(s) per core 值失败: %v\n", err)
				return nil, err
			}
			cpuInfo.Threads = uint32(threads)
		case "Architecture":
			cpuInfo.Architecture = value
		case "Socket(s)":
			cpuNum, err := strconv.ParseUint(value, 10, 32)
			if err != nil {
				fmt.Printf("解析 Socket(s) 值失败: %v\n", err)
				return nil, err
			}
			cpuInfo.CpuNum = uint32(cpuNum)
		}
	}

	return cpuInfo, nil
}

func GetCPUInfo() {
	cmd := exec.CommandContext(context.Background(), "lscpu")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("执行 lscpu 命令失败:", err)
		return
	}

	lines := strings.Split(string(output), "\n")
	var cpuInfo = make(map[string]string)
	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		cpuInfo[key] = value
	}

	fmt.Println("CPU 信息:")
	fmt.Printf("  CPU 数量: %s\n", cpuInfo["CPU(s)"])
	fmt.Printf("  型号名称: %s\n", cpuInfo["Model name"])
	fmt.Printf("  架构: %s\n", cpuInfo["Architecture"])
	fmt.Printf("  核心数量: %s\n", cpuInfo["Core(s) per socket"])
	fmt.Printf("  线程数/核心: %s\n", cpuInfo["Thread(s) per core"])
	fmt.Printf("  插槽数: %s\n", cpuInfo["Socket(s)"])
}
