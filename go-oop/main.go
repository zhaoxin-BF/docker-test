package main

import (
	"fmt"
	"github.com/zhaoxin-BF/docker-test/go-oop/vertex"
)

func main() {
	test := vertex.New(2, 5)
	ret := test.Abs()
	fmt.Println(ret)
	fmt.Println("Hello World", test.X)
}
