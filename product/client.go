package product

import (
	"github.com/medium-stories/go-mono-repo/internal/grpc"
	pbProduct "github.com/medium-stories/go-mono-repo/product/proto"
)

type client struct {
	rpcClient pbProduct.ProductClient
}

func NewClient(accAddr string) *client {
	conn := grpc.CreateClientConnection(accAddr)
	rpcClient := pbProduct.NewProductClient(conn)

	return &client{
		rpcClient: rpcClient,
	}
}
