package repository

import (
	"context"
	"github.com/medium-stories/go-mono-repo/user"
)

type inmemory struct {
	Accounts []*user.Account
}

func NewInMemory() *inmemory {
	return &inmemory{}
}

func (repo *inmemory) AddAccount(ctx context.Context, account *user.Account) (*user.Account, error) {
	return &user.Account{}, nil
}
