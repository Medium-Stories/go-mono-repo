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
	return &user.Account{}, nil
}

func (repo *pgDb) DeleteAccount(ctx context.Context, id string) error {
	return nil
}

func (repo *pgDb) GetById(ctx context.Context, id string) (*user.Account, error) {
	return &user.Account{}, nil
}

func (repo *pgDb) GetByEmail(ctx context.Context, email string) (*user.Account, error) {
	return &user.Account{}, nil
}
