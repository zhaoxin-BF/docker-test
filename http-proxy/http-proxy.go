package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

type Proxy struct {
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get worker
	workerId := r.Header.Get("x-everai-worker-id")
	if workerId == "" {
		http.Error(w, "http header no set worker id", http.StatusBadRequest)
		return
	}
	fmt.Printf("ServeHttp Get Worker Id: %s", workerId)

	a := httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL.Scheme = "http"
			r.URL.Host = "127.0.0.1:8080"
		},
	}

	a.ServeHTTP(w, r)
}

func RunHttp(port int) error {
	addr := fmt.Sprintf(":%d", port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	svc := &http.Server{
		Handler: new(Proxy),
	}
	return svc.Serve(lis)
}

func main() {
	RunHttp(6688)
}
