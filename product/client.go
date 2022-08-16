package product

import (
	"github.com/medium-stories/go-mono-repo/internal/grpc"
	pbProduct "github.com/medium-stories/go-mono-repo/product/proto"
)

func NewClient(accAddr string) pbProduct.ProductClient {
	conn := grpc.CreateClientConnection(accAddr)
	return pbProduct.NewProductClient(conn)
}
