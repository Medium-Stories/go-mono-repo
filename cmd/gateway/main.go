package main

import (
	"flag"
	"github.com/medium-stories/go-mono-repo/gateway"
	"github.com/medium-stories/go-mono-repo/internal/web"
)

var (
	httpAddr    = flag.String("http", ":8000", "Http address")
	accountAddr = flag.String("account_uri", ":8001", "User Account Service address")
)

func main() {
	flag.Parse()

	router := web.NewRouter()

	api := gateway.NewApi(*accountAddr)

	router.POST("users", api.CreateAccount())
	router.DELETE("users/:id", api.DeleteAccount())

	web.ServeHttp(*httpAddr, "gateway", router)
}
