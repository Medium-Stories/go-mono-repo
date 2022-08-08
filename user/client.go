package user

import (
	"github.com/medium-stories/go-mono-repo/internal/grpc"
	pbUser "github.com/medium-stories/go-mono-repo/user/proto"
)

type client struct {
	rpcClient pbUser.AccountManagementClient
}

func NewClient(accAddr string) *client {
	conn := grpc.CreateClientConnection(accAddr)
	rpcClient := pbUser.NewAccountManagementClient(conn)

	return &client{
		rpcClient: rpcClient,
	}
}
