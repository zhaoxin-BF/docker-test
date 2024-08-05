package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			go func() {
				time.Sleep(10 * time.Second)
				fmt.Println("Tick at", t.Format("15:04:05"))
			}()
		}
	}
}
