package utils

import (
	"fmt"
	"os"
	"path"
)

var (
	EverAINodeHome string
	VolumeRoot     string
	DBPath         string
)

func SetEveraiNodeHome(home string) {
	EverAINodeHome = home
	VolumeRoot = GetVolumeRoot()
	DBPath = GetEverAIDbPath()
	EverAINodeHome = GetEverAINodeHome()
}

func GetEverAINodeHome() string {
	if EverAINodeHome != "" {
		return EverAINodeHome
	}
	everaiHome := os.Getenv("EVERAI_NODE_HOME")
	if everaiHome != "" {
		return everaiHome
	}
	cacheDir := os.Getenv("XDG_CACHE_HOME")
	if cacheDir == "" {
		cacheDir = getDefaultDir()
	}
	return path.Join(cacheDir, "everai", "node")
}

func getDefaultDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return path.Join(homeDir, ".cache")
}

func GetVolumeRoot() string {
	root := os.Getenv("EVERAI_VOLUME_ROOT")
	if root != "" {
		return root
	}
	return path.Join(GetEverAINodeHome(), "volumes")
}

func GetEverAIDbPath() string {
	root := os.Getenv("EVERAI_DB_Path")
	if root != "" {
		return root
	}
	CheckAndCreatePath(GetEverAINodeHome())
	return path.Join(GetEverAINodeHome(), "resource_node.db")
}

func CheckAndCreatePath(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Sprintf("Unable to create path: %s\n", err)
			return err
		}
	}
	return nil
}
