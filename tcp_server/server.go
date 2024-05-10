package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}
		if n > 0 {
			request := string(buf[:n])
			fmt.Println("Received:", request)

			response := "world"
			_, err := conn.Write([]byte(response))
			if err != nil {
				fmt.Println("Write error:", err)
				return
			}
			fmt.Println("Sent:", response)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8086")
	if err != nil {
		fmt.Println("Listen error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started, listening on localhost:8086")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			return
		}

		go handleConnection(conn)
	}
}
