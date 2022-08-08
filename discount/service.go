package discount

import (
	"context"
	pbDiscount "github.com/medium-stories/go-mono-repo/discount/proto"
	"github.com/medium-stories/go-mono-repo/internal/grpc"
	pbProduct "github.com/medium-stories/go-mono-repo/product/proto"
	grpcLib "google.golang.org/grpc"
)

const serviceName = "discount service"

type service struct {
	pbDiscount.UnimplementedDiscountProviderServer
	addr string
	repo Repository
}

type Discount struct{}

type Filter struct{}

type Repository interface {
	GetDiscounts(ctx context.Context, filter *Filter) ([]*Discount, error)
}

func NewService(addr string, repo Repository) *service {
	return &service{
		addr: addr,
		repo: repo,
	}
}

func (svc *service) ListenForConnections(ctx context.Context) {
	grpc.ListenForConnections(ctx, svc, svc.addr, serviceName)
}

func (svc *service) RegisterGrpcServer(server *grpcLib.Server) {
	pbDiscount.RegisterDiscountProviderServer(server, svc)
}

func (svc *service) ApplyDiscount(ctx context.Context, req *pbDiscount.DiscountsRequest) (*pbDiscount.DiscountsResponse, error) {
	svc.repo.GetDiscounts(ctx, &Filter{})

	applyDiscount(req.Products, []*Discount{})

	return &pbDiscount.DiscountsResponse{}, nil
}

func applyDiscount(products []*pbProduct.ProductMessage, discounts []*Discount) []*pbProduct.ProductMessage {
	return nil
}
