package server

import (
	"math/rand"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.RegisteredClaims
}

type TransferRequest struct {
	ToAccount int `json:"toAccount"`
	Amount    int `json:"amount"`
}

type createAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firsName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreateAt  time.Time `json:"createAt"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000)),
		CreateAt:  time.Now().UTC(),
	}
}
