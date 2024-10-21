package main

import (
	"fmt"
	"math"
)

//func main() {
//	// 保存当前 umask 值
//	oldMask := syscall.Umask(0)
//	defer syscall.Umask(oldMask)
//
//	// 创建目录并设置权限
//	err := os.MkdirAll("./hello", 0777)
//	if err != nil {
//		fmt.Println("Error creating directory:", err)
//		return
//	}
//
//	// 查看创建的目录权限
//	fileInfo, err := os.Stat("./hello")
//	if err != nil {
//		fmt.Println("Error getting file info:", err)
//		return
//	}
//	fmt.Println("Directory permissions:", fileInfo.Mode())
//
//	wg := sync.WaitGroup{}
//	cha := make(chan int, 1)
//	chb := make(chan int, 1)
//	wg.Add(2)
//	go printA(&wg, cha, chb)
//	go printB(&wg, cha, chb)
//	cha <- 1
//	wg.Wait()
//	close(cha)
//	close(chb)
//}
//
//func printA(wg *sync.WaitGroup, cha, chb chan int) {
//	defer wg.Done()
//	for i := 0; i < 10; i++ {
//		<-cha
//		fmt.Println("A", i)
//		chb <- 1
//	}
//	return
//}
//
//func printB(wg *sync.WaitGroup, cha, chb chan int) {
//	defer wg.Done()
//
//	for i := 0; i < 10; i++ {
//		<-chb
//		fmt.Println("B", i)
//		cha <- 1
//	}
//	return
//}

//func main() {
//	count := 10
//	num := 1000
//	wg := sync.WaitGroup{}
//
//	c := make(chan struct{}, count)
//	for i := 0; i < num; i++ {
//		wg.Add(1)
//		c <- struct{}{}
//		go func(j int) {
//			defer wg.Done()
//			time.Sleep(time.Second * 2)
//			fmt.Println(j)
//			<-c
//		}(i)
//	}
//	wg.Wait()
//}

// 给出一个数字在map数组中找到最相近的数字
func main() {
	ranks := make(map[int]int)
	ho := 31
	ranks[1] = 93
	ranks[10] = 55
	ranks[15] = 30
	ranks[20] = 19
	ranks[23] = 11
	ranks[30] = 2

	gap := 0
	prevGap := 0
	result := 0

	for rank, game := range ranks {
		gap = int(math.Abs(float64(game - ho)))
		fmt.Println(gap)
		if gap <= prevGap {
			result = rank
		}
		prevGap = gap
	}
	fmt.Println(result)
}
