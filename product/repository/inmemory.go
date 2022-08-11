package repository

import (
	"context"
	"github.com/medium-stories/go-mono-repo/product"
)

type inmemory struct {
	products []*productInMemory
}

type productInMemory struct{}

func NewInMemory() product.Repository {
	return &inmemory{}
}

func (repo *inmemory) GetProductsByFilter(ctx context.Context, filter *product.Filter) ([]*product.Product, error) {
	return []*product.Product{}, nil
}
