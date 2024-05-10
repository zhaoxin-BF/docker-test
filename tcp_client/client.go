package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8086")
	if err != nil {
		fmt.Println("Dial error:", err)
		return
	}
	defer conn.Close()

	for i := 0; i < 10; i++ {
		request := "hello"
		_, err = conn.Write([]byte(request))
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}
		fmt.Println("Sent:", request)

		responseBuf := make([]byte, 1024)
		n, err := conn.Read(responseBuf)
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}
		if n > 0 {
			response := string(responseBuf[:n])
			fmt.Println("Received:", response)
		}
		time.Sleep(1 * time.Second)
	}
}
