package model

import (
	"github.com/go-faker/faker/v4"
	"github.com/jackc/pgx/v5/pgtype"
)

type Email struct {
	ID        pgtype.Int8        `json:"id"`
	Campaign  Campaign           `json:"campaign"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	IsActive  pgtype.Bool        `json:"is_active"`
	IsSuccess pgtype.Bool        `json:"is_success"`
	Email     pgtype.Text        `json:"email"`
	FirstName pgtype.Text        `json:"first_name"`
	LastName  pgtype.Text        `json:"last_name"`
	Container map[string]any     `json:"container"`
}

type Emails struct {
	Emails []Email `json:"emails"`
}

func FakeEmail() {
	err := faker.FakeData(&Email{})
	if err != nil {
		panic(err)
	}
}
