package common

//func GetTimezone() string {
//	currentTime := time.Now()
//	zone, offset := currentTime.Zone()
//	location := time.FixedZone(zone, offset)
//	zoneName, _ := time.Now().In(location).Zone()
//	fmt.Println(zoneName)
//
//	// 获取程序运行宿主机的本地时区
//	localZone, err := time.LoadLocation("")
//	if err != nil {
//		fmt.Println("无法获取本地时区:", err)
//		return localZone.String()
//	}
//
//	// 打印结果
//	fmt.Printf("程序运行宿主机的时区：%s\n", localZone.String())
//	return zoneName
//}

import (
	"bytes"
	"fmt"
	"github.com/thlib/go-timezone-local/tzlocal"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp/syntax"
	"runtime"
	"strings"
	"time"
	"unicode/utf8"
	// "golang.org/x/sys/windows/registry"
)

// zoneDirs adapted from https://golang.org/src/time/zoneinfo_unix.go

// https://golang.org/doc/install/source#environment
// list of available GOOS as of 10th Feb 2017
// android, darwin, dragonfly, freebsd, linux, netbsd, openbsd, plan9, solaris,windows

var zoneDirs = map[string]string{
	"android":   "/system/usr/share/zoneinfo/",
	"darwin":    "/usr/share/zoneinfo/",
	"dragonfly": "/usr/share/zoneinfo/",
	"freebsd":   "/usr/share/zoneinfo/",
	"linux":     "/usr/share/zoneinfo/",
	"netbsd":    "/usr/share/zoneinfo/",
	"openbsd":   "/usr/share/zoneinfo/",
	// "plan9":"/adm/timezone/", -- no way to test this platform
	"solaris": "/usr/share/lib/zoneinfo/",
	"windows": `SOFTWARE\Microsoft\Windows NT\CurrentVersion\Time Zones\`,
}

var zoneDir string

var timeZones []string

// InSlice ... check if an element is inside a slice
func InSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

// ReadTZFile ... read timezone file and append into timeZones slice
func ReadTZFile(path string) {
	files, _ := ioutil.ReadDir(zoneDir + path)
	for _, f := range files {
		if f.Name() != strings.ToUpper(f.Name()[:1])+f.Name()[1:] {
			continue
		}
		if f.IsDir() {
			ReadTZFile(path + "/" + f.Name())
		} else {
			tz := (path + "/" + f.Name())[1:]
			// check if tz is already in timeZones slice
			// append if not
			if !InSlice(tz, timeZones) { // need a more efficient method...

				// convert string to rune
				tzRune, _ := utf8.DecodeRuneInString(tz[:1])

				if syntax.IsWordChar(tzRune) { // filter out entry that does not start with A-Za-z such as +VERSION
					timeZones = append(timeZones, tz)
				}
			}
		}
	}

}

func ListTimeZones() {
	if runtime.GOOS == "nacl" || runtime.GOOS == "" {
		fmt.Println("Unsupported platform")
		os.Exit(0)
	}

	// detect OS
	fmt.Println("Time zones available for : ", runtime.GOOS)
	fmt.Println("------------------------")

	fmt.Println("Retrieving time zones from : ", zoneDirs[runtime.GOOS])

	if runtime.GOOS != "windows" {
		for _, zoneDir = range zoneDirs {
			ReadTZFile("")
		}
	} else { // let's handle Windows
		// if you're building this on darwin/linux
		// chances are you will encounter
		// undefined: registry in registry.OpenKey error message
		// uncomment below if compiling on Windows platform

		//k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion\Time Zones`, registry.ENUMERATE_SUB_KEYS|registry.QUERY_VALUE)

		//if err != nil {
		// fmt.Println(err)
		//}
		//defer k.Close()

		//names, err := k.ReadSubKeyNames(-1)
		//if err != nil {
		// fmt.Println(err)
		//}

		//fmt.Println("Number of timezones : ", len(names))
		//for i := 0; i <= len(names)-1; i++ {
		// check if tz is already in timeZones slice
		// append if not
		// if !InSlice(names[i], timeZones) { // need a more efficient method...
		//  timeZones = append(timeZones, names[i])
		// }
		//}

		// UPDATE : Reading from registry is not reliable
		// better to parse output result by "tzutil /g" command
		// REMEMBER : There is no time difference between Coordinated Universal Time and Greenwich Mean Time ....
		cmd := exec.Command("tzutil", "/l")

		data, err := cmd.Output()

		if err != nil {
			panic(err)
		}

		fmt.Println("UTC is the same as GMT")
		fmt.Println("There is no time difference between Coordinated Universal Time and Greenwich Mean Time ....")
		GMTed := bytes.Replace(data, []byte("UTC"), []byte("GMT"), -1)

		fmt.Println(string(GMTed))

	}

	now := time.Now()

	for _, v := range timeZones {

		if runtime.GOOS != "windows" {

			location, err := time.LoadLocation(v)
			if err != nil {
				fmt.Println(err)
			}

			// extract the GMT
			t := now.In(location)
			t1 := fmt.Sprintf("%s", t.Format(time.RFC822Z))
			tArray := strings.Fields(t1)
			gmtTime := strings.Join(tArray[4:], "")
			hours := gmtTime[0:3]
			minutes := gmtTime[3:]

			gmt := "GMT" + fmt.Sprintf("%s:%s", hours, minutes)
			fmt.Println(gmt + " " + v)

		} else {
			fmt.Println(v)
		}

	}
	fmt.Println("Total timezone ids : ", len(timeZones))
}

func GetOSTimezone() {
	// 获取当前操作系统
	os := runtime.GOOS

	// 获取当前时间
	currentTime := time.Now()

	// 获取时区
	var zoneName string
	if os == "windows" {
		zoneName = "Local"
	} else {
		zone, _ := currentTime.Zone()
		zoneName = zone
	}

	// 打印结果
	fmt.Printf("操作系统：%s\n", os)
	fmt.Printf("时区：%s\n", zoneName)
}

func GetTZ() {
	ltz, _ := tzlocal.LocalTZ()
	fmt.Println(ltz)

	etz, _ := tzlocal.EnvTZ()
	fmt.Println(etz)

	rtz, _ := tzlocal.RuntimeTZ()
	fmt.Println(rtz)
}
