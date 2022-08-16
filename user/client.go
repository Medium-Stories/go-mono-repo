package user

import (
	"github.com/medium-stories/go-mono-repo/internal/grpc"
	pbUser "github.com/medium-stories/go-mono-repo/user/proto"
)

func NewClient(accAddr string) pbUser.AccountManagementClient {
	conn := grpc.CreateClientConnection(accAddr)
	return pbUser.NewAccountManagementClient(conn)
}
