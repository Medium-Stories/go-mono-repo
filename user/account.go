package user

import "time"

type Account struct {
	Id        string
	FirstName string
	LastName  string
	Nickname  string
	Password  string
	Email     string
	Country   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
