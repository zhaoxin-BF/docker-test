package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	rootPath := "/Users/zhaoxin/Workspace/github.com/zhaoxin-BF/docker-test/newDir"
	workerPath := path.Join(rootPath, "worker-3")
	workerVolumePath := path.Join(workerPath, "empty", "worker-3-volume")

	err := os.MkdirAll(workerVolumePath, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileInstance, err := os.Create(path.Join(workerVolumePath, "username"))
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = fileInstance.Write([]byte("hello world"))
	fileInstance.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(workerVolumePath)
	//str, err := os.MkdirTemp(workerPath, "myapp-")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(str)

	//time.Sleep(10 * time.Second)
	//err := os.RemoveAll(workerPath)
	//if err != nil {
	//	fmt.Println(err)
	//}
}
