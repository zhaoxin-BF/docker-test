package main

import (
	"fmt"
	"github.com/zhaoxin-BF/docker-test/bindata/plugins/sc_bindata"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	pwdDir := "/Users/zhaoxin/Workspace/github.com/zhaoxin-BF/docker-test/bindata/temp"
	for _, asset := range sc_bindata.AssetNames() {
		target := filepath.Join(pwdDir, asset)
		targetPath := filepath.Dir(target)
		if err := Mkdir(targetPath); err != nil {
			fmt.Println("Mkdir err:", err)
			return
		}
		content, err := sc_bindata.Asset(asset)
		if err != nil {
			return
		}
		if err := WriteFile(target, content); err != nil {
			return
		}
	}

	data, err := sc_bindata.Asset("scripts/check-step.sh")
	if err != nil {
		fmt.Println(err)
		// Asset was not found.
	}
	fmt.Println("------------------------------: ")
	//fmt.Println(string(data))

	// use asset data
	// 创建一个临时文件来存储脚本
	tmpFile, err := ioutil.TempFile("", "script-*.sh")
	if err != nil {
		fmt.Println("Error creating temporary file:", err)
		return
	}
	defer os.Remove(tmpFile.Name())

	// 将脚本内容写入临时文件
	_, err = tmpFile.Write(data)
	if err != nil {
		fmt.Println("Error writing script to file:", err)
		return
	}

	// 设置脚本可执行权限
	err = os.Chmod(tmpFile.Name(), 0755)
	if err != nil {
		fmt.Println("Error setting file permissions:", err)
		return
	}

	// 执行脚本
	time.Sleep(1 * time.Second)
	cmd := exec.Command(tmpFile.Name(), "false")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		//fmt.Println("Error running script:", err)
		return
	}
}

func WriteFile(filename string, content []byte) error {
	err := os.WriteFile(filename, content, 0755)
	return err
}

func Mkdir(path string) error {
	err := os.Mkdir(path, 0755)
	if err != nil && os.IsExist(err) {
		// If the error is that the directory already exists, return nil
		return nil
	}
	return err
}
