package main

import (
	"context"
	"flag"
	"github.com/medium-stories/go-mono-repo/product"
	"github.com/medium-stories/go-mono-repo/product/discount"
	"github.com/medium-stories/go-mono-repo/product/repository"
)

var (
	addr         = flag.String("addr", ":8002", "Product Service address")
	discountsUri = flag.String("discounts_uri", ":8003", "Discounts Service address")
)

func main() {
	flag.Parse()

	svc := product.NewService(*addr, repository.NewInMemory(), &discount.Provider{Addr: *discountsUri})

	rootCtx, rootCancel := context.WithCancel(context.Background())
	defer rootCancel()

	svc.ListenForConnections(rootCtx)
}
