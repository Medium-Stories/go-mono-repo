package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/medium-stories/go-mono-repo/user"
	"time"
)

type inmemory struct {
	Accounts []*user.Account
}

func NewAccountInmemory() *inmemory {
	return &inmemory{}
}

func (repo *inmemory) AddAccount(ctx context.Context, account *user.Account) (*user.Account, error) {
	account.Id = uuid.New().String()
	account.CreatedAt = time.Now().UTC()
	account.UpdatedAt = time.Now().UTC()

	repo.Accounts = append(repo.Accounts, account)

	return account, nil
}

func (repo *inmemory) DeleteAccount(ctx context.Context, id string) error {
	for i, acc := range repo.Accounts {
		if acc.Id == id {
			copy(repo.Accounts[i:], repo.Accounts[i+1:])
			repo.Accounts[len(repo.Accounts)-1] = nil
			repo.Accounts = repo.Accounts[:len(repo.Accounts)-1]
			break
		}
	}

	return nil
}

func (repo *inmemory) GetById(ctx context.Context, id string) (*user.Account, error) {
	return repo.getById(id), nil
}

func (repo *inmemory) GetByEmail(ctx context.Context, email string) (*user.Account, error) {
	return repo.getByEmail(email), nil
}

func (repo *inmemory) getById(id string) *user.Account {
	for _, acc := range repo.Accounts {
		if acc.Id == id {
			return acc
		}
	}

	return nil
}

func (repo *inmemory) getByEmail(email string) *user.Account {
	for _, acc := range repo.Accounts {
		if acc.Email == email {
			return acc
		}
	}

	return nil
}
