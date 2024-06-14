package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	rootPath := "/Users/zhaoxin/Workspace/github.com/zhaoxin-BF/docker-test/newDir"
	workerPath := path.Join(rootPath, "worker-3")

	//err := os.MkdirAll(path.Join(rootPath, "worker-3", "empty"), 0777)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//str, err := os.MkdirTemp(workerPath, "myapp-")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(str)

	//time.Sleep(10 * time.Second)
	err := os.RemoveAll(workerPath)
	if err != nil {
		fmt.Println(err)
	}
}
