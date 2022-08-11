package repository

import (
	"context"
	"github.com/medium-stories/go-mono-repo/discount"
)

type inmemory struct {
	discounts []*discountInMemory
}

type discountInMemory struct{}

func NewInMemory() discount.Repository {
	return &inmemory{}
}

func (repo *inmemory) GetDiscounts(ctx context.Context, filter *discount.Filter) ([]*discount.Discount, error) {
	return []*discount.Discount{}, nil
}
