package device_resource

import (
	"errors"
	"fmt"
	"github.com/gyuho/linux-inspect/df"
	"os/exec"
	"strings"
)

type FilesystemInfo struct {
	// Device name
	Device string `json:"device"`
	// UUID
	Uuid string `json:"uuid"`
	// File system type (e.g., ext4, btrfs)
	Type                  string `json:"type"`
	TotalBlocksBytesN     int64  `json:"total_blocks_bytes_n"`
	AvailableBlocksBytesN int64  `json:"available_blocks_bytes_n"`
	UsedBlocksBytesN      int64  `json:"used_blocks_bytes_n"`
	// Percentage of filesystem used
	UsedPercent string `json:"usedPercent"`
	// Mounted filesystem path
	Mountpoint string `json:"mountpoint"`
}

type FilesystemCollector struct {
	Filesystem FilesystemInfo
}

func NewFilesystemCollector() *FilesystemCollector {
	return &FilesystemCollector{}
}

func (f *FilesystemCollector) GetFilesystemInfo() (fileSystemInfo []FilesystemInfo, err error) {
	// df pkg not support other os.
	// TODO only ubuntu
	infos, err := df.GetDefault("")
	//fmt.Println("----------------info:", infos)
	if err != nil {
		return nil, err
	}

	for _, info := range infos {
		uuid := ""
		if info.MountedOn == "/" {
			uuid, _ = f.getRootFilesystemUuid(info.FileSystem)
		}
		var temp FilesystemInfo
		temp.Device = info.Device
		temp.Uuid = uuid
		temp.Type = info.FileSystemType
		temp.TotalBlocksBytesN = info.TotalBlocksBytesN
		temp.AvailableBlocksBytesN = info.AvailableBlocksBytesN
		temp.UsedBlocksBytesN = info.UsedBlocksBytesN
		temp.UsedPercent = info.IusedPercent
		temp.Mountpoint = info.MountedOn

		fileSystemInfo = append(fileSystemInfo, temp)
		if info.MountedOn == "/" {
			fmt.Printf("%+v\n", temp)
		}
	}

	return
}

func (f *FilesystemCollector) getRootFilesystemUuid(device string) (string, error) {
	out, err := exec.Command("blkid").Output()
	if err != nil {
		return "", nil
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		fields := strings.Split(line, ":")
		if fields[0] == device {
			fmt.Println(line)
			return getUuid(line)
		}
	}
	return "", errors.New("not found uuid")
}

func getUuid(str string) (uuid string, err error) {
	// get UUID="**"
	fields := strings.Split(str, " ")
	// get **
	for _, val := range fields {

		if strings.Contains(val, "UUID=") {
			fmt.Printf("%+v\n", val)
			uuids := strings.Split(val, "\"")
			uuid = uuids[1]
			return uuid, err
		}
	}
	return "", errors.New("not found uuid")
}
