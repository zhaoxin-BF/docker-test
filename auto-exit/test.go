package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	n := 9
	if n < 8 || n > 32 {
		n = 8
	}
	fullUUID := uuid.New().String()
	fmt.Println("Full UUID:", fullUUID)

	shortUUID := uuid.New().String()[:n]
	fmt.Println("Short UUID:", shortUUID)
}
