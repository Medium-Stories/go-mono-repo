package main

import (
	"flag"
	"github.com/medium-stories/go-mono-repo/internal/web"
	"github.com/medium-stories/go-mono-repo/product"
	"github.com/medium-stories/go-mono-repo/user"
)

var (
	httpAddr    = flag.String("http", ":8000", "Http address")
	accountAddr = flag.String("account_uri", ":8001", "User Account Service address")
	productsUri = flag.String("products_uri", ":8002", "Products Service address")
)

func main() {
	flag.Parse()

	router := web.NewRouter()

	userClient := user.NewClient(*accountAddr)
	productsClient := product.NewClient(*productsUri)

	router.POST("users", user.CreateAccountHandler(userClient))
	router.GET("products", product.GetProductsByFilterHandler(productsClient))

	web.ServeHttp(*httpAddr, "gateway", router)
}
