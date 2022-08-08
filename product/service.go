package product

import (
	"context"
	"github.com/medium-stories/go-mono-repo/internal/grpc"
	pbProduct "github.com/medium-stories/go-mono-repo/product/proto"
	grpcLib "google.golang.org/grpc"
)

const serviceName = "product service"

type service struct {
	pbProduct.UnimplementedProductServer
	addr             string
	repo             Repository
	discountProvider DiscountProvider
}

type Product struct{}

type Filter struct{}

type Repository interface {
	GetProductsByFilter(ctx context.Context, filter *Filter) ([]*Product, error)
}

// DiscountProvider communicates to discount service. It is responsible for discounts on products.
type DiscountProvider interface {
	ApplyDiscount(ctx context.Context, products []*pbProduct.ProductMessage) ([]*pbProduct.ProductMessage, error)
}

func NewService(addr string, repo Repository, discountProvider DiscountProvider) *service {
	return &service{
		addr:             addr,
		repo:             repo,
		discountProvider: discountProvider,
	}
}

func (svc *service) ListenForConnections(ctx context.Context) {
	grpc.ListenForConnections(ctx, svc, svc.addr, serviceName)
}

func (svc *service) RegisterGrpcServer(server *grpcLib.Server) {
	pbProduct.RegisterProductServer(server, svc)
}

func (svc *service) GetProductsByFilter(ctx context.Context, req *pbProduct.GetProductsByFilterRequest) (*pbProduct.ProductsResponse, error) {
	svc.repo.GetProductsByFilter(ctx, &Filter{})

	svc.discountProvider.ApplyDiscount(ctx, []*pbProduct.ProductMessage{})

	return &pbProduct.ProductsResponse{}, nil
}
