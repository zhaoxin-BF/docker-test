package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"os"
)

func ConatinerExec() {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.43"))
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	// 获取容器 ID
	containerID := "ef44da5a7878"

	// 在容器内执行脚本
	execConfig := types.ExecConfig{
		//Cmd: []string{"echo 'ready' > /hellos.txt"},
		Cmd:          []string{"/bin/sh", "-c", "echo 'ready' > /hellos.txt"},
		AttachStdout: true,
		AttachStderr: true,
	}
	exec, err := cli.ContainerExecCreate(context.Background(), containerID, execConfig)
	if err != nil {
		panic(err)
	}

	attach, err := cli.ContainerExecAttach(context.Background(), exec.ID, types.ExecStartCheck{})
	if err != nil {
		panic(err)
	}
	defer attach.Close()

	// 读取执行结果
	_, err = io.Copy(os.Stdout, attach.Reader)
	if err != nil {
		panic(err)
	}

	// 检查执行状态
	inspect, err := cli.ContainerExecInspect(context.Background(), exec.ID)
	if err != nil {
		panic(err)
	}
	if inspect.ExitCode != 0 {
		fmt.Printf("Script execution failed with exit code %d\n", inspect.ExitCode)
	} else {
		fmt.Println("Script executed successfully")
	}
}
