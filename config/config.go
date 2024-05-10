package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Base Base `yaml:"base"`
}

type Base struct {
	DBPath     string `yaml:"db_path"`
	VolumePath string `yaml:"volume_path"`
}

func GetConfig() {
	// 读取YAML文件内容
	yamlFile, err := ioutil.ReadFile("/Users/zhaoxin/Workspace/github.com/zhaoxin-BF/everai-test/config/config.yaml")
	if err != nil {
		panic(err)
	}

	// 定义一个Config类型的变量
	var config Config

	// 解析YAML文件
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	// 打印配置信息
	fmt.Printf("Config: %+v\n", config)
}
