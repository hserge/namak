package service

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Email struct {
	ID        pgtype.Int4        `json:"id"`
	Campaign  Campaign           `json:"campaign"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	IsActive  pgtype.Bool        `json:"is_active"`
	IsSuccess pgtype.Bool        `json:"is_success"`
	Email     pgtype.Text        `json:"email"`
	FirstName pgtype.Text        `json:"first_name"`
	LastName  pgtype.Text        `json:"last_name"`
	Container pgtype.JSONCodec   `json:"container"`
}

type EmailService struct {
	ctx context.Context
	db  *pgx.Conn
}

func NewEmailService(ctx context.Context, db *pgx.Conn) *EmailService {
	return &EmailService{
		ctx: ctx,
		db:  db,
	}
}

func (es *EmailService) All() ([]Email, error) {
	rows, err := es.db.Query(es.ctx, "SELECT id, created_at, is_active, is_success, email, first_name, last_name, container FROM emails order by id desc")
	if err != nil {
		return []Email{}, err
	}
	defer rows.Close()

	var result []Email
	for rows.Next() {
		email := Email{}
		err := rows.Scan(&email.ID, &email.CreatedAt, &email.IsActive, &email.IsSuccess, &email.Email, &email.FirstName, &email.LastName, &email.Container)
		// Exit if we get an error
		if err != nil {
			return []Email{}, err
		}
		result = append(result, email)
	}
	return result, nil
}

func (es *EmailService) FindByID(ID int) (*Email, error) {
	return &Email{}, nil
}
