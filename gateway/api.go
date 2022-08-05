package gateway

import (
	"github.com/medium-stories/go-mono-repo/internal/grpc"
	pbUser "github.com/medium-stories/go-mono-repo/user/proto"
)

type api struct {
	rpcClient pbUser.AccountManagementClient
}

func NewApi(accAddr string) *api {
	conn := grpc.CreateClientConnection(accAddr)
	client := pbUser.NewAccountManagementClient(conn)

	return &api{
		rpcClient: client,
	}
}
