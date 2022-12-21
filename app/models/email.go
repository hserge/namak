package models

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Email struct {
	Id        pgtype.Int4        `json:"id"`
	Campaign  Campaign           `json:"campaign"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	IsActive  pgtype.Bool        `json:"is_active"`
	IsSuccess pgtype.Bool        `json:"is_success"`
	Email     pgtype.Text        `json:"email"`
	FirstName pgtype.Text        `json:"first_name"`
	LastName  pgtype.Text        `json:"last_name"`
	Container pgtype.JSONCodec   `json:"container"`
}

type EmailRepository struct {
	db *pgx.Conn
}

func NewEmailRepository(db *pgx.Conn) *EmailRepository {
	return &EmailRepository{
		db,
	}
}

func (r *EmailRepository) All() ([]Email, error) {
	rows, err := r.db.Query(context.Background(), "SELECT id, created_at, is_active, is_success, email, first_name, last_name, container FROM emails order by id desc")
	if err != nil {
		return []Email{}, err
	}
	defer rows.Close()

	var result []Email
	for rows.Next() {
		email := Email{}
		err := rows.Scan(&email.Id, &email.CreatedAt, &email.IsActive, &email.IsSuccess, &email.Email, &email.FirstName, &email.LastName, &email.Container)
		// Exit if we get an error
		if err != nil {
			return []Email{}, err
		}
		result = append(result, email)
	}
	return result, nil
}

func (r *EmailRepository) Get(id int) (Email, error) {
	return Email{}, nil
}

func (r *EmailRepository) Update(db *sql.DB) (Email, error) {
	return Email{}, errors.New("not implemented")
}

func (r *EmailRepository) Delete(db *sql.DB) (bool, error) {
	return true, errors.New("not implemented")
}

func (r *EmailRepository) Create(db *sql.DB) (Email, error) {
	return Email{}, errors.New("not implemented")
}
