package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// 定义源目录和目标目录
	volume1 := "combine"
	volume2 := "combine1"
	combine := "combine-volume"

	// 创建 combine 目录
	if err := os.MkdirAll(combine, os.ModePerm); err != nil {
		fmt.Println("创建 combine 目录失败:", err)
		return
	}

	// 处理 volume1
	if err := linkFiles(volume1, combine); err != nil {
		fmt.Println("链接 volume1 中的文件失败:", err)
	}

	// 处理 volume2
	if err := linkFiles(volume2, combine); err != nil {
		fmt.Println("链接 volume2 中的文件失败:", err)
	}
}

// 生成随机前缀
func generateRandomPrefix() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%d_", rand.Intn(100000)) // 生成一个 0-9999 的随机数作为前缀
}

func linkFiles(sourceDir, targetDir string) error {
	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 计算链接的目标路径
		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}
		linkPath := filepath.Join(targetDir, relPath)

		// 如果是目录，创建目录
		if info.IsDir() {
			return os.MkdirAll(linkPath, os.ModePerm)
		}

		// 检查链接路径是否存在
		if _, err := os.Lstat(linkPath); err == nil {
			// 目标路径已存在，添加随机前缀
			prefix := generateRandomPrefix()
			linkPath = replaceFileName(linkPath, prefix)
		}

		// 如果是文件，创建符号链接
		return os.Link(path, linkPath)
	})
}

// 替换文件名函数
func replaceFileName(filePath, newPrefix string) string {
	dir := filepath.Dir(filePath)          // 获取目录部分
	base := filepath.Base(filePath)        // 获取文件名部分
	newFileName := newPrefix + "." + base  // 生成新的文件名
	return filepath.Join(dir, newFileName) // 返回新的完整路径
}
