package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"strings"
	"time"
)

type CollectorMetrics struct {
	Name             string
	ID               string
	CPUPercentage    float64
	Memory           int64
	MemoryLimit      int64
	MemoryPercentage float64
	NetworkRx        int64
	NetworkTx        int64
	BlockRead        float64
	BlockWrite       float64
	PidsCurrent      uint64
	Timestamp        time.Time
}

func ContainerStats(Id string) {
	apiClient, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	containerID := Id

	containerStats, err := apiClient.ContainerStats(context.Background(), containerID, false)
	if err != nil {
		panic(err)
	}

	defer containerStats.Body.Close()

	var v *types.StatsJSON
	dec := json.NewDecoder(containerStats.Body)
	if err := dec.Decode(&v); err != nil {
	}

	fmt.Printf("%+v\n", v.CPUStats)

	info := getCollectorMetrics(v)
	fmt.Printf("%+v\n", info)
}

// 处理types.StatsJSON数据
func getCollectorMetrics(stats *types.StatsJSON) *CollectorMetrics {
	fmt.Printf("%+v\n", stats)
	var (
		memPercent        = 0.0
		cpuPercent        = 0.0
		blkRead, blkWrite uint64
		mem               = 0.0
		memLimit          = 0.0
		pids              uint64
		netRx             = 0.0
		netTx             = 0.0
	)

	//cpu
	fmt.Println("stats.CPUStats.CPUUsage.TotalUsage", stats.CPUStats.CPUUsage.TotalUsage)
	fmt.Println("stats.PreCPUStats.CPUUsage.TotalUsage", stats.PreCPUStats.CPUUsage.TotalUsage)
	cpuDelta := float64(stats.CPUStats.CPUUsage.TotalUsage - stats.PreCPUStats.CPUUsage.TotalUsage)
	fmt.Println("cpuDelta: ", cpuDelta)
	systemDelta := float64(stats.CPUStats.SystemUsage - stats.PreCPUStats.SystemUsage)
	fmt.Println("systemDelta: ", systemDelta)
	cpuPercent = (cpuDelta / systemDelta) * 100.0 * float64(stats.CPUStats.OnlineCPUs)

	//memory
	mem = float64(stats.MemoryStats.Usage)
	memLimit = float64(stats.MemoryStats.Limit)
	if stats.MemoryStats.Limit != 0 {
		fmt.Println("memoryLimit: ", stats.MemoryStats.Limit)
		fmt.Println("memoryUsage: ", stats.MemoryStats.Usage)
		memPercent = float64(stats.MemoryStats.Usage) / float64(stats.MemoryStats.Limit) * 100.0
	}
	//network
	for _, v := range stats.Networks {
		netRx += float64(v.RxBytes)
		netTx += float64(v.TxBytes)
	}
	//block
	var blkio = stats.BlkioStats
	for _, bioEntry := range blkio.IoServiceBytesRecursive {
		switch strings.ToLower(bioEntry.Op) {
		case "read":
			blkRead = blkRead + bioEntry.Value
		case "write":
			blkWrite = blkWrite + bioEntry.Value
		}
	}
	//pidsCurrent
	pids = stats.PidsStats.Current
	return &CollectorMetrics{
		Name:             stats.Name,
		ID:               stats.ID,
		CPUPercentage:    cpuPercent,
		Memory:           int64(mem),
		MemoryLimit:      int64(memLimit),
		MemoryPercentage: memPercent,
		NetworkRx:        int64(netRx),
		NetworkTx:        int64(netTx),
		BlockRead:        float64(blkRead),
		BlockWrite:       float64(blkWrite),
		PidsCurrent:      pids,
		Timestamp:        stats.Read,
	}
}
