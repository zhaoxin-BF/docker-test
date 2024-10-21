package main

import (
	"github.com/joho/godotenv"
	"os"
)

func main() {
	//ConatinerExec()
	GetContainerLogsPro()
	//GetcContainer()
	//StopContainer()
	//Load()
	//everaiHome := os.Getenv("EVERAI_NODE_HOME")
	//fmt.Println("EVERAI_NODE_HOME:", everaiHome)

}

func Load() error {
	_, err := os.Stat(".env")
	if err == nil {
		return godotenv.Load()
	}

	if os.IsNotExist(err) {
		return nil
	}
	return err
}
