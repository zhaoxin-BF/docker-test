package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func StopContainer() {
	apiClient, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	containerID := "b5a79628e2745a65d3be08282f25f335f5c68d82eddf04d7b678f3d1071462f8s"

	err = apiClient.ContainerStop(context.Background(), containerID, container.StopOptions{})
	//fmt.Println(topList)
	if err != nil {
		fmt.Println(err)
		return
	}
}
