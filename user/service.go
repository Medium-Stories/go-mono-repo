package user

import (
	"context"
	"github.com/medium-stories/go-mono-repo/event"
	"github.com/medium-stories/go-mono-repo/internal/grpc"
	pbUser "github.com/medium-stories/go-mono-repo/user/proto"
	grpcLib "google.golang.org/grpc"
)

const serviceName = "account management service"

// accountService will expose account management service via grpc
type accountService struct {
	pbUser.UnimplementedAccountManagementServer
	addr    string
	repo    AccountRepository
	pub     AccountPublisher
	pwdHash PasswordHash
}

// You can implement these interfaces anywhere you want, in any package.
// You could create /repository, /publisher as subpackages in /cmd/user/
// and keep user service completely clean of any implementations. Or,
// you could create such subpackages in user service and it all together.
// Both approaches have their own pros and cons, it's your choice to pick the one you like.

// AccountRepository communicates to data store with user accounts
type AccountRepository interface {
	AddAccount(ctx context.Context, account *Account) (*Account, error)
	DeleteAccount(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) (*Account, error)
	GetByEmail(ctx context.Context, email string) (*Account, error)
}

// AccountPublisher will publish event that corresponds to an account action
type AccountPublisher interface {
	Publish(event string, msg interface{}) error
}

// PasswordHash is used to hash and validate user password
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
	svc.pwdHash.Hash(req.Password)

	svc.repo.AddAccount(ctx, &Account{})

	svc.pub.Publish(event.AccountCreated, 1)

	return &pbUser.AccountMessage{}, nil
}

func (svc *accountService) DeleteAccount(ctx context.Context, req *pbUser.DeleteAccountRequest) (*pbUser.DeleteAccountResponse, error) {
	svc.repo.DeleteAccount(ctx, req.Id)

	svc.pub.Publish(event.AccountDeleted, 1)

	return &pbUser.DeleteAccountResponse{}, nil
}
