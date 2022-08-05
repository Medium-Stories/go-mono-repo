package events

import (
	"context"
	"github.com/gobackpack/rmq"
	"github.com/medium-stories/go-mono-repo/event"
)

type accountDeleted struct {
	hub *rmq.Hub
}

func NewAccountDeletedListener(hub *rmq.Hub) *accountDeleted {
	return &accountDeleted{
		hub: hub,
	}
}

func (ev *accountDeleted) Listen(ctx context.Context) {
	consumer := startConsumer(ctx, ev.hub, event.AccountDeleted)
	go handleMessages(ctx, consumer, event.AccountDeleted)
}
