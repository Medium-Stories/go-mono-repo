package events

import (
	"context"
	"github.com/gobackpack/rmq"
	"github.com/sirupsen/logrus"
	"time"
)

// Listener will listen for account events
type Listener interface {
	Listen(ctx context.Context)
}

func Listen(ctx context.Context, listeners ...Listener) {
	for _, listener := range listeners {
		listener.Listen(ctx)
	}
}

// startConsumer will start rmq consumer
func startConsumer(ctx context.Context, hub *rmq.Hub, event string) *rmq.Consumer {
	conf := rmq.NewConfig()
	conf.Exchange = "account"
	conf.Queue = event
	conf.RoutingKey = event

	if err := hub.CreateQueue(conf); err != nil {
		logrus.Fatal(err)
	}

	return hub.StartConsumer(ctx, conf)
}

// handleMessages is default message handler, it will only log received messages
func handleMessages(ctx context.Context, cons *rmq.Consumer, name string) {
	logrus.Infof("%s started", name)

	defer logrus.Warnf("%s closed", name)

	for {
		select {
		case msg := <-cons.OnMessage:
			logrus.Infof("[%s] %s - %s", time.Now().UTC(), name, msg)
		case err := <-cons.OnError:
			logrus.Error(err)
		case <-ctx.Done():
			return
		}
	}
}
