package device_resource_new

//
//import (
//	"context"
//	"fmt"
//	"github.com/shirou/gopsutil/cpu"
//	"os/exec"
//	"strings"
//)
//
//type CpuInfo struct {
//	Model        string // CPU model, for example: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
//	Cores        uint32 // CPU core, for example: CPU(s): 12
//	Threads      uint32 // CPU Thread(s) per core，for example：Thread(s) per core:  2
//	Architecture string // CPU arch, for example: x86_64/arm64/~
//	CpuNum       uint32 // Number of cpus, In most cases, there will only be one physical cpu
//}
//
//type CpuCollector struct {
//	CPU CpuInfo
//}
//
//func NewCPUCollector() *CpuCollector {
//	return &CpuCollector{}
//}
//
//func (c *CpuCollector) GetCPUInfo() (*CpuInfo, error) {
//	var cpuInfo *CpuInfo
//	// exec.Command
//	cmd := exec.Command("lscpu")
//	stdout, err := cmd.CombinedOutput()
//	if err != nil {
//		fmt.Println("执行 lscpu 命令失败:", err)
//		return &c.CPU, nil
//	}
//	result := strings.Split(string(stdout), "\n")
//	//fmt.Printf("%+v", result)
//	for _, line := range result {
//		fmt.Printf("%+v\n", line)
//		if strings.HasPrefix(line, "Model name:") {
//			cpuInfo.Model = strings.TrimSpace(strings.Split(line, ":")[1])
//		} else if strings.HasPrefix(line, "CPU(s):") {
//			cpuInfo.Cores = parseUint32(strings.TrimSpace(strings.Split(line, ":")[1]))
//		} else if strings.HasPrefix(line, "Thread(s) per core:") {
//			cpuInfo.Threads = parseUint32(strings.TrimSpace(strings.Split(line, ":")[1]))
//		} else if strings.HasPrefix(line, "Architecture:") {
//			cpuInfo.Architecture = strings.TrimSpace(strings.Split(line, ":")[1])
//		} else if strings.HasPrefix(line, "Socket(s):") {
//			cpuInfo.CpuNum = parseUint32(strings.TrimSpace(strings.Split(line, ":")[1]))
//		}
//	}
//
//	return &c.CPU, nil
//}
//
//func parseUint32(s string) uint32 {
//	var value uint32
//	fmt.Sscanf(s, "%d", &value)
//	return value
//}
//func GetCPUInfo() {
//	info, err := cpu.InfoWithContext(context.Background())
//	if err != nil {
//		return
//	}
//
//	// test printf
//	fmt.Println(info)
//
//	for _, val := range info {
//		fmt.Println("CPU 信息:")
//		fmt.Printf("  CPU 数量: %d\n", val.CPU)
//		fmt.Printf("  厂商 ID: %s\n", val.VendorID)
//		fmt.Printf("  系列: %s\n", val.Family)
//		fmt.Printf("  型号: %s\n", val.Model)
//		fmt.Printf("  Stepping: %d\n", val.Stepping)
//		fmt.Printf("  物理 ID: %s\n", val.PhysicalID)
//		fmt.Printf("  核心 ID: %s\n", val.CoreID)
//		fmt.Printf("  核心数量: %d\n", val.Cores)
//		fmt.Printf("  型号名称: %s\n", val.ModelName)
//		fmt.Printf("  主频: %.2f MHz\n", val.Mhz)
//		fmt.Printf("  缓存大小: %d KB\n", val.CacheSize)
//		fmt.Printf("  Flags: %v\n", val.Flags)
//		fmt.Printf("  微码: %s\n", val.Microcode)
//	}
//}
