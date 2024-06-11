package container

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"io"
	"os"
	"strings"
	"time"
)

func GetContainerLogs() {
	apiClient, err := client.NewClientWithOpts(client.WithVersion("1.43"))
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	out, err := apiClient.ContainerLogs(context.Background(), "54e43b3949514988dec1cf72c3bb4341f8175808117c91a05644ba166a860862", container.LogsOptions{ShowStdout: true, Follow: true})
	if err != nil {
		panic(err)
	}
	defer out.Close()

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}

func GetStreamLogs() {
	apiClient, err := client.NewClientWithOpts(client.WithVersion("1.43"))
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	containerID := "021079273as79bb454adde43c4a2bbaf2b1344a7930ef5c89631ce977dee0ae61"

	out, err := apiClient.ContainerLogs(context.Background(), containerID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		//Timestamps: true,
		Since: "2024-05-09T06:39:09.0Z",
		//Until:      "2024-05-09T06:39:15.0Z",
		//Tail: "172.17.0.1 - - [09/May/2024:06:39:13 +0000] \"GET / HTTP/1.0\" 200 615 \"-\" \"ApacheBench/2.3\" \"-\"",
	})
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// 读取并打印数据帧
	readAndPrintFrames(out)
}

func readAndPrintFrames(src io.ReadCloser) (err error) {
	buf := make([]byte, 512)
	nr := 0

	go CloseOut(src)

	for {
		nr, err = src.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read response body failed")
				return
			}
		}
		if nr > 0 {
			wStr := buf[0:nr]
			// send response
			//fmt.Printf("%+v\n", string(wStr))
			lines := strings.Split(string(wStr), "\n")
			for _, line := range lines {
				fmt.Printf("%+v\n", line)
			}
		}
	}
}

func CloseOut(src io.ReadCloser) {
	time.Sleep(10000 * time.Second)
	fmt.Println("close container stdout")
	src.Close()
	fmt.Println("close container stderr")
	return
}

//

func readAndPrintFramesPro(src io.ReadCloser) (err error) {
	buf := make([]byte, 32*1024)
	nr := 0

	timeout := time.After(5 * time.Minute) // set timeout
ReadLoop:
	for {
		select {
		case <-timeout:
			fmt.Println("Timeout reached, closing src and exiting program")
			src.Close()
			break ReadLoop
		default:
			nr, err = src.Read(buf)
			if err != nil {
				if err != io.EOF {
					fmt.Println("read response body failed")
					break ReadLoop
				}
				break ReadLoop
			}
			if nr > 0 {
				wStr := buf[0:nr]
				// send response
				fmt.Print(string(wStr))
			}
		}
	}
	return err
}
