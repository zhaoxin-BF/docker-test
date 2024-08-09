package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

func main() {
	//volumeCtx, cancel := context.WithCancel(context.Background())
	//
	//for i := 0; i < 2; i++ {
	//	go func() {
	//		//defer fmt.Println("Goroutine exited")
	//		for j := 0; j < 10; j++ {
	//			select {
	//			case <-volumeCtx.Done():
	//				fmt.Println("Goroutine received cancel signal")
	//				return
	//			default:
	//				fmt.Println("Goroutine received signal")
	//			}
	//		}
	//	}()
	//}
	//
	//time.Sleep(10 * time.Second)
	//fmt.Println("Canceling context")
	//cancel()
	//
	//time.Sleep(5 * time.Second)
	//fmt.Println("Done")

	data1 := []interface{}{"apple", "banana", "cherry", "durian"}
	data2 := []interface{}{"red", "yellow", "green", "brown"}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Column 1", "Column 2"})

	for i := 0; i < len(data1); i++ {
		table.Append([]string{fmt.Sprintf("%v", data1[i]), fmt.Sprintf("%v", data2[i])})
	}

}
