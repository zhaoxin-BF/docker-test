package setup

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func SyncFiles(srcPath string, descPath string) error {
	// 检查源路径是否存在
	srcInfo, err := os.Stat(srcPath)
	if err != nil {
		return fmt.Errorf("Error accessing source path: %v", err)
	}

	// 如果源路径是文件,则直接复制文件
	if !srcInfo.IsDir() {
		return CopyFile(srcPath, descPath)
	}

	// 如果源路径是目录,则递归复制目录下的所有文件
	return filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcPath, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(descPath, relPath)

		if info.IsDir() {
			return os.MkdirAll(targetPath, srcInfo.Mode())
		} else {
			return CopyFile(path, targetPath)
		}
	})
}

func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("Error opening source file: %v", err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("Error creating destination file: %v", err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("Error copying file: %v", err)
	}
	return nil
}
