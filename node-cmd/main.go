package main

import (
	"context"
	"github.com/zhaoxin-BF/docker-test/node-cmd/cmd"
	"log"
)

var (
	PROGRAM = "everai-resource-node"
	VERSION = "devel"
)

func main() {
	ctx := context.WithValue(context.TODO(), "VERSION", VERSION)
	ctx = context.WithValue(ctx, "PROGRAM", PROGRAM)

	if err := cmd.Execute(ctx); err != nil {
		log.Fatal(err)
	}
}
