package discount

import (
	"context"
	pbDiscount "github.com/medium-stories/go-mono-repo/discount/proto"
	"github.com/medium-stories/go-mono-repo/internal/grpc"
	pbProduct "github.com/medium-stories/go-mono-repo/product/proto"
)

type Provider struct {
	Addr string
}

func (provider *Provider) ApplyDiscount(ctx context.Context, products []*pbProduct.ProductMessage) ([]*pbProduct.ProductMessage, error) {
	conn := grpc.CreateClientConnection(provider.Addr)
	discountsClient := pbDiscount.NewDiscountProviderClient(conn)

	resp, err := discountsClient.ApplyDiscount(ctx, &pbDiscount.DiscountsRequest{
		Products: products,
	})

	return resp.Products, err
}
