package main

import (
	"context"
	"flag"
	"github.com/medium-stories/go-mono-repo/discount"
	"github.com/medium-stories/go-mono-repo/discount/repository"
)

var (
	addr = flag.String("addr", ":8003", "Discounts Service address")
)

func main() {
	flag.Parse()

	svc := discount.NewService(*addr, repository.NewInMemory())

	rootCtx, rootCancel := context.WithCancel(context.Background())
	defer rootCancel()

	svc.ListenForConnections(rootCtx)
}
