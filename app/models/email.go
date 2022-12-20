package models

import (
	"database/sql"
	"errors"

	"github.com/hserge/namak/app"
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

type Emails struct {
	Emails []Email `json:"emails"`
}

type App app.App

func (a *App) list() (Emails, error) {
	rows, err := a.Db.Query(a.Ctx, "SELECT id, created_at, is_active, is_success, email, first_name, last_name, container FROM emails order by id desc")
	if err != nil {
		return Emails{}, err
	}
	defer rows.Close()

	result := Emails{}
	for rows.Next() {
		email := Email{}
		err := rows.Scan(&email.Id, &email.CreatedAt, &email.IsActive, &email.IsSuccess, &email.Email, &email.FirstName, &email.LastName, &email.Container)
		// Exit if we get an error
		if err != nil {
			return Emails{}, err
		}
		result.Emails = append(result.Emails, email)
	}
	return result, nil
}

func (a *App) get(id int) (Email, error) {
	return Email{}, nil
}

func (e *Email) update(db *sql.DB) (Email, error) {
	return Email{}, errors.New("not implemented")
}

func (e *Email) delete(db *sql.DB) (bool, error) {
	return true, errors.New("not implemented")
}

func (e *Email) create(db *sql.DB) (Email, error) {
	return Email{}, errors.New("not implemented")
}

func getEmails(db *sql.DB, start, count int) ([]Email, error) {
	return nil, errors.New("not implemented")
}
