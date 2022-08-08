package main

import (
	"context"
	"flag"
	"github.com/medium-stories/go-mono-repo/cmd/product/discount"
	"github.com/medium-stories/go-mono-repo/cmd/product/repository"
	"github.com/medium-stories/go-mono-repo/product"
)

var (
	addr         = flag.String("addr", ":8002", "Product Service address")
	discountsUri = flag.String("discounts_uri", ":8003", "Discounts Service address")
)

func main() {
	flag.Parse()

	svc := product.NewService(*addr, repository.NewInMemoryRepository(), &discount.Provider{Addr: *discountsUri})

	rootCtx, rootCancel := context.WithCancel(context.Background())
	defer rootCancel()

	svc.ListenForConnections(rootCtx)
}
