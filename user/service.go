package user

import (
	"context"
	"errors"
	"github.com/medium-stories/go-mono-repo/event"
	"github.com/medium-stories/go-mono-repo/internal/grpc"
	pbUser "github.com/medium-stories/go-mono-repo/user/proto"
	"github.com/sirupsen/logrus"
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

// Filter to apply when querying data store for user accounts
type Filter struct {
	Page    int
	Limit   int
	Country string
}

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

// AddAccount will add new user account
func (svc *accountService) AddAccount(ctx context.Context, req *pbUser.AccountRequest) (*pbUser.AccountMessage, error) {
	existing, err := svc.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if existing != nil && existing.Email == req.Email {
		return nil, errors.New("email already exists")
	}

	hashed, err := svc.pwdHash.Hash(req.Password)
	if err != nil {
		return nil, err
	}

	req.Password = hashed

	account, err := svc.repo.AddAccount(ctx, protoReqToUserAccount(req))
	if err != nil {
		return nil, err
	}

	go func(id string) {
		if pubErr := svc.pub.Publish(event.AccountCreated, id); pubErr != nil {
			logrus.Error(pubErr)
		}
	}(account.Id)

	return userAccountToProto(account), nil
}

func (svc *accountService) DeleteAccount(ctx context.Context, req *pbUser.DeleteAccountRequest) (*pbUser.DeleteAccountResponse, error) {
	account, err := svc.repo.GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if account == nil || account.Email == "" {
		return nil, errors.New("account not found")
	}

	if err = svc.repo.DeleteAccount(ctx, req.Id); err != nil {
		return nil, err
	}

	go func(id string) {
		if pubErr := svc.pub.Publish(event.AccountDeleted, id); pubErr != nil {
			logrus.Error(pubErr)
		}
	}(req.Id)

	return &pbUser.DeleteAccountResponse{
		Success: true,
	}, nil
}

func userAccountToProto(account *Account) *pbUser.AccountMessage {
	return &pbUser.AccountMessage{
		Id:        account.Id,
		FirstName: account.FirstName,
		LastName:  account.LastName,
		Nickname:  account.Nickname,
		Password:  account.Password,
		Email:     account.Email,
		Country:   account.Country,
		CreatedAt: account.CreatedAt.String(),
		UpdatedAt: account.UpdatedAt.String(),
		DeletedAt: account.DeletedAt.String(),
	}
}

func protoReqToUserAccount(pbAccount *pbUser.AccountRequest) *Account {
	return &Account{
		FirstName: pbAccount.FirstName,
		LastName:  pbAccount.LastName,
		Nickname:  pbAccount.Nickname,
		Password:  pbAccount.Password,
		Email:     pbAccount.Email,
		Country:   pbAccount.Country,
	}
}
