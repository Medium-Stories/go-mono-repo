package repository

import (
	"context"
	"github.com/medium-stories/go-mono-repo/product"
)

type productInMemory struct{}

// inMemoryRepository is in memory data store implementation for products
type inMemoryRepository struct {
	products []*productInMemory
}

func NewInMemoryRepository() product.Repository {
	return &inMemoryRepository{}
}

func (repo *inMemoryRepository) GetProductsByFilter(ctx context.Context, filter *product.Filter) ([]*product.Product, error) {
	return []*product.Product{}, nil
}
