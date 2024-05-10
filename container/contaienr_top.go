package container

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
	"strconv"
	"strings"
)

func ContainerTop() {
	apiClient, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	containerID := "b5a79628e2745a65d3be08282f25f335f5c68d82eddf04d7b678f3d1071462f8"

	topList, err := apiClient.ContainerTop(context.Background(), containerID, nil)
	//fmt.Println(topList)
	if err != nil {
		fmt.Println(err)
		return
	}
	pidIndex := -1
	fmt.Printf("top list: %v\n", topList.Titles)
	for k, v := range topList.Titles {
		if strings.ToLower(v) == "pid" {
			pidIndex = k
		}
	}
	if pidIndex < 0 {
		fmt.Println("Not found pid")
		return
	}
	fmt.Printf("top pid: %v\n", topList.Processes)
	for _, v := range topList.Processes {
		if len(v) < pidIndex {
			fmt.Println("Not found pid")
			continue
		}
		val, err := strconv.Atoi(v[pidIndex])
		if err != nil {
			fmt.Println("Not found pid")
			continue
		}
		fmt.Println(val)
		fmt.Println(containerID)
	}
}
