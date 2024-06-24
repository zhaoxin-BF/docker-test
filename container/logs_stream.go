package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func readJSONLogs(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()

	scanner := bufio.NewScanner(file)
	offset := int64(0)

	for {
		// 移动文件指针到上次读取的位置
		_, err = file.Seek(offset, 0)
		if err != nil {
			log.Println("Error seeking file:", err)
		}

		for scanner.Scan() {
			line := scanner.Bytes()

			var data map[string]interface{}
			err := json.Unmarshal(line, &data)
			if err != nil {
				log.Println("Error parsing JSON:", err)
				continue
			}

			// 在这里对每行数据进行处理
			// 例如，打印数据
			fmt.Println(data)
		}

		if err := scanner.Err(); err != nil {
			return err
		}

		// 获取当前文件大小
		fileInfo, err := file.Stat()
		if err != nil {
			return err
		}
		currentSize := fileInfo.Size()

		// 更新偏移量为当前文件大小
		offset = currentSize

		// 如果文件大小没有变化，说明没有新的日志写入，阻塞等待
		if currentSize == fileSize {
			time.Sleep(1 * time.Second)
			continue
		}

		// 更新文件大小为当前文件大小
		fileSize = currentSize
	}
}

func StreamRead() {
	// 替换为你的JSON日志文件路径
	filePath := "./test.json"
	err := readJSONLogs(filePath)
	if err != nil {
		log.Fatal(err)
	}
}
