package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {
	// 生成一个 100 MB 的随机数据文件
	generateRandomFile("test_file.txt", 100*1024*1024)

	// 测试 32 KB 缓冲区
	start := time.Now()
	copyWithBuffer("test_file.txt", "output_32kb.txt", 32*1024)
	fmt.Printf("Copy with 32 KB buffer: %s\n", time.Since(start))

	// 测试 1 MB 缓冲区
	start = time.Now()
	copyWithBuffer("test_file.txt", "output_1mb.txt", 1024*1024)
	fmt.Printf("Copy with 1 MB buffer: %s\n", time.Since(start))
}

func generateRandomFile(filename string, size int64) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// 生成随机数据并写入文件
	data := make([]byte, 4096)
	rand.Seed(time.Now().UnixNano())
	for i := int64(0); i < size/4096; i++ {
		rand.Read(data)
		_, err = file.Write(data)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

func copyWithBuffer(src, dst string, bufferSize int) {
	sourceFile, err := os.Open(src)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer destFile.Close()

	buf := make([]byte, bufferSize)
	_, err = io.CopyBuffer(destFile, sourceFile, buf)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}
}
