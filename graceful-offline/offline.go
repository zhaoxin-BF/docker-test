package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/gracefulOffline", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			message := r.URL.Query().Get("action")
			if message == "exit" {
				fmt.Fprintln(w, "Goodbye!")
				time.Sleep(20 * time.Second)
				os.Exit(0)
			}
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
