package user

import (
	"context"
	"github.com/medium-stories/go-mono-repo/event"
	"github.com/medium-stories/go-mono-repo/internal/grpc"
	pbUser "github.com/medium-stories/go-mono-repo/user/proto"
	grpcLib "google.golang.org/grpc"
)

const serviceName = "account management service"

type accountService struct {
	pbUser.UnimplementedAccountManagementServer
	addr    string
	repo    AccountRepository
	pub     AccountPublisher
	pwdHash PasswordHash
}

type Account struct{}

type AccountRepository interface {
	AddAccount(ctx context.Context, account *Account) (*Account, error)
}

type AccountPublisher interface {
	Publish(event string, msg interface{}) error
}

type PasswordHash interface {
	Hash(plain string) (string, error)
	Validate(hashed, plain string) bool
}

func NewAccountService(addr string, repo AccountRepository, pub AccountPublisher, pwdHash PasswordHash) *accountService {
	return &accountService{
		addr:    addr,
		repo:    repo,
		pub:     pub,
		pwdHash: pwdHash,
	}
}

func (svc *accountService) ListenForConnections(ctx context.Context) {
	grpc.ListenForConnections(ctx, svc, svc.addr, serviceName)
}

func (svc *accountService) RegisterGrpcServer(server *grpcLib.Server) {
	pbUser.RegisterAccountManagementServer(server, svc)
}

func (svc *accountService) AddAccount(ctx context.Context, req *pbUser.AccountRequest) (*pbUser.AccountMessage, error) {
	svc.pwdHash.Hash("12345")

	svc.repo.AddAccount(ctx, &Account{})

	svc.pub.Publish(event.AccountCreated, 1)

	return &pbUser.AccountMessage{}, nil
}
