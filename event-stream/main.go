package main

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"
)

func testStreamClient() {
	http.HandleFunc("/stream2", func(writer http.ResponseWriter, request *http.Request) {
		cli := http.Client{}
		req, err := http.NewRequest("GET", "http://127.0.0.1:9091/stream", nil)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}
		resp, err := cli.Do(req)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}
		fmt.Printf("header: %+v\n", resp.Header)
		fmt.Printf("status code:%v\n", resp.Status)
		for k, v := range resp.Header {
			for _, vv := range v {
				writer.Header().Add(k, vv)
			}
		}
		defer resp.Body.Close()
		writerFlusher := writer.(http.Flusher)
		buf := make([]byte, 1)
		for {
			nr, er := resp.Body.Read(buf)
			if nr > 0 {
				wStr := buf[0:nr]
				fmt.Println("wStr: ", string(wStr))
				writer.WriteHeader(http.StatusOK)
				_, ew := writer.Write(wStr)
				if ew != nil {
					err = ew
					break
				}
				writerFlusher.Flush()
				time.Sleep(2 * time.Second)
			}
			if er != nil {
				fmt.Println("writer error:", er)
				if er != io.EOF {
					err = er
				}
				break
			}
		}
		if err != nil {
			fmt.Println("read err: ", err.Error())
		}

		fmt.Printf("end\n")
	})

	err := http.ListenAndServe(":9092", nil)
	if err != nil {
		fmt.Println("listen failed :", err.Error())
	} else {
		fmt.Println("end")
	}
}

func TestHttpStream(t *testing.T) {
	go testStreamClient()
	http.HandleFunc("/stream", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "text/event-stream")
		fl, ok := writer.(http.Flusher)
		if !ok {
			fmt.Println("http.ResponseWriter not type http.Flusher")
			return
		}
		//for i := 0; i < 1000; i++ {
		fmt.Fprint(writer, fmt.Sprintf("message %s\n", "sfsfdsfwsdfdssdddddddfsdfsfsfsfsdfs"))

		fl.Flush()
		//time.Sleep(1 * time.Second)
		//}

		writer.WriteHeader(http.StatusAccepted)

	})
	fmt.Println("start server")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("err: ", err.Error())
	} else {
		fmt.Println("end")
	}
}

func main() {
	TestHttpStream(&testing.T{})
}
