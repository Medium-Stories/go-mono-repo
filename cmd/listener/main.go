package main

import (
	"context"
	"flag"
	"github.com/gobackpack/rmq"
	"github.com/medium-stories/go-mono-repo/cmd/listener/events"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

var rmqHost = flag.String("rmq_host", "localhost", "RabbitMQ host address")

func main() {
	flag.Parse()

	cred := rmq.NewCredentials()
	cred.Host = *rmqHost
	hub := rmq.NewHub(cred)

	hubCtx, hubCancel := context.WithCancel(context.Background())
	defer hubCancel()

	_, err := hub.Connect(hubCtx)
	if err != nil {
		logrus.Fatal(err)
	}

	// create listeners for different account actions/events
	accountCreated := events.NewAccountCreatedListener(hub)
	accountDeleted := events.NewAccountDeletedListener(hub)

	events.Listen(hubCtx, accountCreated, accountDeleted)

	logrus.Info("listening for messages...")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
