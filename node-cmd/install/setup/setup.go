package setup

import (
	"fmt"
	"os"
)

func SetUp() {
	pwdDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("------ print dir ------", pwdDir)

	//sFiles := path.Join(pwdDir, "everai_resource_node.service")
	//descFiles := path.Join(pwdDir, "everai_resource_node.service.desc")
	//
	//SyncFiles(sFiles, descFiles)
	//
	//sFiles = path.Join(pwdDir, "install")
	//descFiles = path.Join(pwdDir, "comm", "install")
	//
	//SyncFiles(sFiles, descFiles)
}
