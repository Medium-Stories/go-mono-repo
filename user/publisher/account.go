package publisher

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gobackpack/rmq"
	"github.com/medium-stories/go-mono-repo/event"
	"github.com/sirupsen/logrus"
)

type accountPublisher struct {
	hub  *rmq.Hub
	conf map[string]*rmq.Publisher
}

func NewAccountPublisher(ctx context.Context, hub *rmq.Hub) *accountPublisher {
	pub := &accountPublisher{
		hub:  hub,
		conf: make(map[string]*rmq.Publisher),
	}

	pub.setupEvents(ctx, []string{event.AccountCreated})

	return pub
}

func (pub *accountPublisher) Publish(event string, msg interface{}) error {
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	rmqPub := pub.conf[event]
	if rmqPub == nil {
		return errors.New("invalid account event")
	}

	rmqPub.Publish(b)

	return nil
}

func (pub *accountPublisher) setupEvents(ctx context.Context, events []string) {
	for _, ev := range events {
		conf := rmq.NewConfig()
		conf.Exchange = "account"
		conf.Queue = ev
		conf.RoutingKey = ev

		if err := pub.hub.CreateQueue(conf); err != nil {
			logrus.Fatal(err)
		}

		pub.conf[ev] = pub.hub.CreatePublisher(ctx, conf)
	}
}
