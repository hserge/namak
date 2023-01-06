package service

import (
	"context"
	"time"

	"github.com/hserge/namak/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmailService struct {
	ctx    context.Context
	dbPool *pgxpool.Pool
}

func NewEmailService(ctx context.Context, dbPool *pgxpool.Pool) *EmailService {
	return &EmailService{ctx: ctx, dbPool: dbPool}
}

func (es *EmailService) Create(email *model.Email) error {
	ctx, cancel := context.WithTimeout(es.ctx, 15*time.Second)
	defer cancel()
	return es.dbPool.QueryRow(ctx, `
        INSERT INTO emails (is_active, email, first_name, last_name, container) 
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at, is_active, email, first_name, last_name, container`,
		email.ID, email.CreatedAt, email.IsActive, email.Email, email.FirstName, email.LastName, email.Container).
		Scan(&email.ID, &email.CreatedAt, &email.IsActive, &email.Email, &email.FirstName, &email.LastName, &email.Container)
}

func (es *EmailService) Get(id int) (email model.Email, err error) {
	ctx, cancel := context.WithTimeout(es.ctx, 15*time.Second)
	defer cancel()
	err = es.dbPool.QueryRow(ctx, `
        SELECT id, created_at, is_active, is_success, email, first_name, last_name, container FROM emails
        WHERE id=$1`,
		id).
		Scan(&email.ID, &email.CreatedAt, &email.IsActive, &email.IsSuccess, &email.Email, &email.FirstName, &email.LastName, &email.Container)

	return email, err
}
