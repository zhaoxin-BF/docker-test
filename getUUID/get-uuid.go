package main

import (
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

const uuidFilePath = "/Users/zhaoxin/Workspace/github.com/zhaoxin-BF/docker-test"

func main() {
	// 1. 查询路径 /data/everai/node/.uuid 文件是否存在
	uuidFile := path.Join(uuidFilePath, ".uuid")
	fmt.Println(uuidFile)
	_, err := os.Stat(uuidFile)
	if err == nil {
		uuidBytes, err := ioutil.ReadFile(uuidFile)
		if err == nil && uuidBytes != nil && len(uuidBytes) >= 36 {
			uuidBytes = uuidBytes[:36]
			uuid := string(uuidBytes)
			fmt.Println(uuid)
			return
		} else {
			fmt.Printf("UUID file %s does not exist %+v", uuidFile, err)
		}
	}

	err = os.MkdirAll(filepath.Dir(uuidFile), 0755)
	if err != nil {
		fmt.Printf("Error creating UUID directory: %s", err)
		return
	}

	newUUID := uuid.New().String()
	newUUIDBytes := []byte(newUUID)
	err = ioutil.WriteFile(uuidFile, newUUIDBytes, 0644)
	if err != nil {
		fmt.Printf("Error creating UUID file: %s", err)
		return
	}
	return
}
