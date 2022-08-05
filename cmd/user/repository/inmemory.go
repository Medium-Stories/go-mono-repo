package repository

import (
	"context"
	"github.com/medium-stories/go-mono-repo/user"
)

type inmemory struct {
	Accounts []*user.Account
}

func NewAccountInmemory() *inmemory {
	return &inmemory{}
}

func (repo *inmemory) AddAccount(ctx context.Context, account *user.Account) (*user.Account, error) {
	return &user.Account{}, nil
}

func (repo *inmemory) DeleteAccount(ctx context.Context, id string) error {
	return nil
}

func (repo *inmemory) GetById(ctx context.Context, id string) (*user.Account, error) {
	return &user.Account{}, nil
}

func (repo *inmemory) GetByEmail(ctx context.Context, email string) (*user.Account, error) {
	return &user.Account{}, nil
}
