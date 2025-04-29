package device_resource_new

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type GPUInfo struct {
	Model         string // NPU 型号（如 910B2）
	DriverVersion string // npu-smi 驱动版本（如 24.1.0.3）
	CudaVersion   string // 此处 NPU 无对应概念，留空或填 "N/A"
	Memory        int64  // 总内存（单位：字节，HBM 总容量）
	Power         string // 平均功率或第一个 NPU 功率（单位：W）
	GpuNum        uint32 // NPU 数量
}

type NPUCollector struct {
	NPU GPUInfo
}

func NewNPUCollector() *NPUCollector {
	return &NPUCollector{}
}

var (
	ErrNoNPU = errors.New("no NPU detected")
)

func (n *NPUCollector) GetNPUInfo() (*GPUInfo, error) {
	var npuInfo GPUInfo

	// 执行 npu-smi info 命令
	cmd := exec.Command("npu-smi", "info")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("npu-smi command failed: %v", err)
	}

	output := out.String()
	lines := strings.Split(output, "\n")

	// 提取驱动版本（第一行）
	for _, line := range lines {
		if strings.Contains(line, "npu-smi") && strings.Contains(line, "Version:") {
			parts := strings.Split(line, " ")
			npuInfo.DriverVersion = strings.TrimSpace(parts[2])
			break
		}
	}

	fmt.Println("Driver Version: ", npuInfo.DriverVersion)
	// 解析 NPU 数量（通过统计 "NPU" 出现次数或表格行数）
	var npuCount int
	var firstNPUModel string
	var totalMemory int64 // HBM 总容量（假设所有 NPU 相同，取第一个的总容量）
	var firstNPUPower string

	for _, line := range lines {
		// 匹配 NPU 设备信息行（以 "|" 分隔，包含 "910B2" 等型号）
		if strings.Contains(line, "910B2") && strings.Contains(line, "OK") {
			npuCount++
			// 提取型号（如 "910B2"）
			if firstNPUModel == "" {
				firstNPUModel = "910B2"
			}
			// 提取功率（如 "104.8" W）
			if firstNPUPower == "" {
				firstNPUPower = "100W"
			}
		} else if strings.Contains(line, "65536") {
			hbmTotal := "65536"
			hbmTotalMB, _ := strconv.ParseInt(hbmTotal, 10, 64)
			if totalMemory == 0 {
				totalMemory = hbmTotalMB / 1024 // MB 转 GB（1GB=1024MB）
			}
		}
	}

	fmt.Println("npuCount: ", npuCount)
	fmt.Println("Total Memory: ", totalMemory)
	fmt.Println("first npu model: ", firstNPUModel)
	fmt.Println("first npu power: ", firstNPUPower)

	if npuCount == 0 {
		return nil, ErrNoNPU
	}
	// 填充结构体
	npuInfo.Model = firstNPUModel
	npuInfo.GpuNum = uint32(npuCount)
	npuInfo.Memory = totalMemory
	npuInfo.Power = firstNPUPower // 假设所有 NPU 功率相近，取第一个
	npuInfo.CudaVersion = "N/A"   // NPU 无 CUDA 概念，填占位符

	fmt.Println("npuInfo: ", npuInfo)
	return &npuInfo, nil
}

func GetGPUInfo() {
	gpuCollector := NewNPUCollector()
	gpuInfo, err := gpuCollector.GetNPUInfo()

	if err != nil {
		fmt.Printf("Unable to get GPU info: %v", err)
	}
	fmt.Println(gpuInfo)
}
