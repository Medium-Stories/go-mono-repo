package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/medium-stories/go-mono-repo/user"
	"gorm.io/gorm"
	"time"
)

type Account struct {
	Id        uuid.UUID `gorm:"primarykey"`
	Firstname string
	Lastname  string
	Nickname  string
	Password  string
	Email     string
	Country   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type pgDb struct {
	db *gorm.DB
}

func NewPgDb(db *gorm.DB) *pgDb {
	db.AutoMigrate(&Account{})

	return &pgDb{
		db: db,
	}
}

func (repo *pgDb) AddAccount(ctx context.Context, account *user.Account) (*user.Account, error) {
	acc := accountToEntity(account)
	acc.Id = uuid.New()

	if err := repo.db.Create(&acc).Error; err != nil {
		return nil, err
	}

	account.Id = acc.Id.String()
	account.CreatedAt = acc.CreatedAt
	account.UpdatedAt = acc.UpdatedAt

	return account, nil
}

func (repo *pgDb) DeleteAccount(ctx context.Context, id string) error {
	var acc *Account
	if err := repo.db.Where("id = ?", id).Find(&acc).Error; err != nil {
		return err
	}

	return repo.db.Delete(acc).Error
}

func (repo *pgDb) GetById(ctx context.Context, id string) (*user.Account, error) {
	var acc *Account
	if err := repo.db.Where("id = ?", id).Find(&acc).Error; err != nil {
		return nil, err
	}
	return entityToAccount(acc), nil
}

func (repo *pgDb) GetByEmail(ctx context.Context, email string) (*user.Account, error) {
	var acc *Account
	if err := repo.db.Where("email = ?", email).Find(&acc).Error; err != nil {
		return nil, err
	}
	return entityToAccount(acc), nil
}

func accountToEntity(account *user.Account) *Account {
	return &Account{
		Firstname: account.FirstName,
		Lastname:  account.LastName,
		Nickname:  account.Nickname,
		Password:  account.Password,
		Email:     account.Email,
		Country:   account.Country,
	}
}

func entityToAccount(acc *Account) *user.Account {
	return &user.Account{
		Id:        acc.Id.String(),
		FirstName: acc.Firstname,
		LastName:  acc.Lastname,
		Nickname:  acc.Nickname,
		Password:  acc.Password,
		Email:     acc.Email,
		Country:   acc.Country,
		CreatedAt: acc.CreatedAt,
		UpdatedAt: acc.UpdatedAt,
		DeletedAt: acc.DeletedAt.Time,
	}
}
